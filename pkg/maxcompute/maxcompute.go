package maxcompute

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

const (
	ACCESS_KEY_SECRET_FIELD_NAME = "accessKeySecret"
)

// Make sure Datasource implements required interfaces. This is important to do
// since otherwise we will only get a not implemented error response from plugin in
// runtime. In this example datasource instance implements backend.QueryDataHandler,
// backend.CheckHealthHandler interfaces. Plugin should not implement all these
// interfaces - only those which are required for a particular task.
var (
	_ backend.QueryDataHandler   = (*MaxComputeDatasource)(nil)
	_ backend.CheckHealthHandler = (*MaxComputeDatasource)(nil)
)

type MaxComputeDatasource struct {
	config MaxComputeDatasourceConfig
}

type MaxComputeDatasourceConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	Project         string
}

func NewDatasource(_ context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	log.DefaultLogger.Info("Creating MaxCompute instance")
	var config MaxComputeDatasourceConfig

	err := json.Unmarshal(settings.JSONData, &config)
	if err != nil {
		log.DefaultLogger.Error("unmarshal mc config", "err", err)
		return nil, fmt.Errorf("unmarshal mc config: %v", err)
	}

	accessKeySecret, ok := settings.DecryptedSecureJSONData[ACCESS_KEY_SECRET_FIELD_NAME]
	if !ok {
		return nil, fmt.Errorf("sk not provided")
	}

	config.AccessKeySecret = accessKeySecret
	return &MaxComputeDatasource{
		config: config,
	}, nil
}
