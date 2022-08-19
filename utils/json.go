package utils

import (
	"encoding/json"
)

// struct -> bytes -> json
func ToJson(data *struct{}) string {
	bytes, err := json.Marshal(data)
	PanicErr(err)
	json := string(bytes)
	return json
}
