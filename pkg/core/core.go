package core

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	PublicKey  string `json:"public_key"`
	AllowedIps string `json:"allowed_ips"`
	//Endpoint   string `json:"endpoint"`
}
