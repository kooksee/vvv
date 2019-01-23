package models

import "encoding/json"

type RpcCall struct {
	Service  string `json:"service"`
	Data     string `json:"data"`
	Callback string `json:"callback"`
}

func (t *RpcCall) Encode() string {
	dt, _ := json.Marshal(t)
	return string(dt)
}

func (t *RpcCall) Decode(data string) error {
	return json.Unmarshal([]byte(data), t)
}
