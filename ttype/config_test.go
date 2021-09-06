package ttype

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFuck(t *testing.T) {
	a := InfluxDB{
		Bucket: "1",
		Org:    "2",
		Token:  "3",
		Href:   "4",
	}
	d, e := json.Marshal(a)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(d))
}
