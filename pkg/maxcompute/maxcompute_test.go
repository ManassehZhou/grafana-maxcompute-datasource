package maxcompute

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"gotest.tools/assert"
)

func TestQueryData(t *testing.T) {
	ds := MaxComputeDatasource{
		config: MaxComputeDatasourceConfig{
			AccessKeyId:     os.Getenv("ALIBABACLOUD_ACCESS_KEY_ID"),
			AccessKeySecret: os.Getenv("ALIBABACLOUD_ACCESS_KEY_SECRET"),
			Endpoint:        os.Getenv("MAXCOMPUTE_ENDPOINT"),
			Project:         os.Getenv("MAXCOMPUTE_PROJECT"),
		},
	}

	resp, err := ds.QueryData(
		context.Background(),
		&backend.QueryDataRequest{
			Queries: []backend.DataQuery{
				{RefID: "A", JSON: mustQueryJSON(t, `select 
				100000000000L as bi,
				3.14159261E+7 as doub,
				"abc" as str,
				DATETIME'2017-11-11 00:00:00' as dt,
				True as bool;`)},
				{RefID: "B", JSON: mustQueryJSON(t, "SELECT 2F;")},
			},
		},
	)
	if err != nil {
		t.Error(err)
	}

	if len(resp.Responses) != 2 {
		t.Fatal("QueryData must return a response")
	}

	respA := resp.Responses["A"]

	if err != nil {
		t.Error(respA.Error)
	}
	frame := respA.Frames[0]

	assert.Equal(t, "bi", frame.Fields[0].Name)
	assert.Equal(t, "doub", frame.Fields[1].Name)
	assert.Equal(t, "str", frame.Fields[2].Name)
	assert.Equal(t, "dt", frame.Fields[3].Name)
	assert.Equal(t, "bool", frame.Fields[4].Name)

	for _, f := range frame.Fields {
		assert.Equal(t, 1, f.Len())
	}
}

func mustQueryJSON(t *testing.T, sql string) []byte {
	t.Helper()

	b, err := json.Marshal(queryRequest{
		RawQueryText: sql,
		QueryType:    "table",
	})
	if err != nil {
		panic(err)
	}
	return b
}
