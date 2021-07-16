package config

import "time"

type JWT struct {
	Secret string
	Issuer string
	Expire time.Duration
}
