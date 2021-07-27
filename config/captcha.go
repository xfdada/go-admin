package config

import "time"

type Captcha struct {
	UseRedis bool
	Height   int
	Width    int
	Length   int
	MaxSkew  float64
	DotCount int
	Expiration	time.Duration
	PreKey	string

}
