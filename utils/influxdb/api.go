package influxdb

import (
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func GetApi() api.WriteAPI {
	return Client.WriteAPI(Config.Org, Config.Bucket)
}

func Store(name string, tags map[string]string, value float64) {
	writeApi := GetApi()
	writeApi.WritePoint(write.NewPoint(name, tags, map[string]interface{}{"store": value}, time.Now()))
	writeApi.Flush()
}
