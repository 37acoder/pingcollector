package influxresultcollect

import (
	"context"
	"fmt"

	"github.com/37acoder/pingcollector/ttype"
	"github.com/37acoder/pingcollector/utils/influxdb"
)

const MetricsName = "ping"

type InfluxResultCollector struct {
}

func (i *InfluxResultCollector) ReturnResponse(ctx context.Context, result ttype.PingResult) {
	fmt.Println("influx store response", result)
	influxdb.Store(MetricsName, map[string]string{
		"domain": result.Domain,
	}, float64(result.PingDelay))
}
