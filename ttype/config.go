package ttype

type Config struct {
	Influx InfluxDB `json:"influx"`
}
type InfluxDB struct {
	Bucket string `json:"bucket,omitempty"`
	Org    string `json:"org,omitempty"`
	Token  string `json:"token,omitempty"`
	Href   string `json:"href,omitempty"`
}
