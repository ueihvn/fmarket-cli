package fmarketclient

import "time"

type Config struct {
	host    string
	retry   int
	timeout time.Duration
}

type Option func(*Config)

func defaultConfig() Config {
	return Config{
		host:    "https://api.fmarket.vn",
		retry:   1,
		timeout: 5 * time.Second,
	}
}
