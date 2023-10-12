package maxcompute

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/data/sqlutil"
)

// queryRequest is an inbound query request as part of a batch of queries sent
// to [(*MaxComputeDatasource).QueryData].
type queryRequest struct {
	RawQueryText string
	QueryType    string
}

// executeResult is an envelope for concurrent query responses.
type executeResult struct {
	refID        string
	dataResponse backend.DataResponse
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (mcds *MaxComputeDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	var (
		wg             sync.WaitGroup
		response       = backend.NewQueryDataResponse()
		executeResults = make(chan executeResult, len(req.Queries))
	)

	for _, dataQuery := range req.Queries {
		query, err := decodeQueryRequest(dataQuery)
		if err != nil {
			response.Responses[dataQuery.RefID] = backend.ErrDataResponse(backend.StatusBadRequest, err.Error())
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			executeResults <- executeResult{
				refID:        query.RefID,
				dataResponse: mcds.query(ctx, *query),
			}
		}()
	}

	wg.Wait()
	close(executeResults)
	for r := range executeResults {
		response.Responses[r.refID] = r.dataResponse
	}

	return response, nil
}

func decodeQueryRequest(dataQuery backend.DataQuery) (*sqlutil.Query, error) {

	var req queryRequest
	if err := json.Unmarshal(dataQuery.JSON, &req); err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	var format sqlutil.FormatQueryOption
	switch req.QueryType {
	case "table":
		format = sqlutil.FormatOptionTable
	case "timeseries":
		format = sqlutil.FormatOptionTimeSeries
	default:
		format = sqlutil.FormatOptionTable
	}

	query := &sqlutil.Query{
		RawSQL:        req.RawQueryText,
		RefID:         dataQuery.RefID,
		MaxDataPoints: dataQuery.MaxDataPoints,
		Interval:      dataQuery.Interval,
		TimeRange:     dataQuery.TimeRange,
		Format:        format,
	}

	// Process macros and execute the query.
	sql, err := sqlutil.Interpolate(query, macros)
	if err != nil {
		return nil, fmt.Errorf("macro interpolation: %w", err)
	}
	query.RawSQL = sql

	return query, nil
}

func (mcds *MaxComputeDatasource) query(ctx context.Context, query sqlutil.Query) backend.DataResponse {
	var resp backend.DataResponse

	conn, err := mcds.getConnection()
	if err != nil {
		return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("connect to mc: %v", err))
	}

	rows, err := conn.Query(query.RawSQL)
	if err != nil {
		return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("query mc: %v", err))
	}

	frame, err := sqlutil.FrameFromRows(rows, -1)
	if err != nil && !errors.Is(err, io.EOF) {
		return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("convert data to frame: %v", err))
	}

	resp.Frames = data.Frames{frame}
	return resp
}
