package main

import (
	"os"

	"github.com/ManassehZhou/maxcompute-datasource/pkg/maxcompute"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func main() {
	if err := datasource.Manage("manassehzhou-maxcompute-datasource", maxcompute.NewDatasource, datasource.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
