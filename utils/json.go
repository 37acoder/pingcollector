package utils

import "encoding/json"

func ToJson(v interface{}) []byte {
	d, _ := json.Marshal(v)
	return d
}
