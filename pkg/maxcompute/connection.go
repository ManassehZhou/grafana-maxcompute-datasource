package maxcompute

import (
	"database/sql"

	"github.com/aliyun/aliyun-odps-go-sdk/odps"
	_ "github.com/aliyun/aliyun-odps-go-sdk/sqldriver"
)

func (mcds *MaxComputeDatasource) getConnection() (*sql.DB, error) {

	config := odps.NewConfig()
	config.AccessId = mcds.config.AccessKeyId
	config.AccessKey = mcds.config.AccessKeySecret
	config.Endpoint = mcds.config.Endpoint
	config.ProjectName = mcds.config.Project

	return sql.Open("odps", config.FormatDsn())
}
