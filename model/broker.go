package model

type Broker struct {
	Key          string `json:"key"`
	Id           int32  `json:"broker_id"`
	Name         string `json:"name"`
	RegisterTime string `json:"register_time"`
	Address      string `json:"address"`
}
