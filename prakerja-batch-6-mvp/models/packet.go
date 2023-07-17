package models

type Packet struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Data    int    `json:"data"`
	SMS     int    `json:"sms"`
	Telepon string `json:"telepon"`
}
