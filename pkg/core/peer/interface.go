package peer

import "time"

type data struct {
	AllowedIps        []string  `json:"allowed_ips"`
	LastHandshakeTime time.Time `json:"last_handshake_time"`
	Endpoint          string    `json:"endpoint"`
	PresharedKey      string    `json:"preshared_key"`
	PublicKey         string    `json:"public_key"`
	ReceiveBytes      int64     `json:"receive_bytes"`
	TransmitBytes     int64     `json:"transmit_bytes"`
}

type inputAdd struct {
	PublicKey  string   `json:"public_key"`
	AllowedIps []string `json:"allowed_ips"`
	Endpoint   string   `json:"endpoint"`
}

type inputDelete struct {
	PublicKey string `json:"public_key"`
}
