package main

import (
	"context"
	"sync"

	"github.com/37acoder/pingcollector/service"
	"github.com/37acoder/pingcollector/service/influxresultcollect"
	"github.com/37acoder/pingcollector/utils"
	"github.com/37acoder/pingcollector/utils/influxdb"
)

func main() {
	config, err := utils.LoadConfigFromJsonFile("./config.json")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	influxdb.Init(config.Influx)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	constSource := service.ConstDomainSource{}
	//resultCollector := service.PrintPingResultReceiver{}
	//resultCollector := inmemory.NewResultCollector()
	resultCollector := &influxresultcollect.InfluxResultCollector{}
	pingService := service.NewDefaultPingService()
	err = pingService.StartPingService(ctx, constSource, resultCollector)
	if err != nil {
		return
	}
	wg.Wait()
}
