package influxdb

import (
	"github.com/37acoder/pingcollector/ttype"
	"github.com/influxdata/influxdb-client-go/v2"
)

var Client influxdb2.Client
var Config ttype.InfluxDB

func Init(db ttype.InfluxDB) {
	// You can generate a Token from the "Tokens Tab" in the UI
	Config = db
	Client = influxdb2.NewClient(db.Href, Config.Token)
}
