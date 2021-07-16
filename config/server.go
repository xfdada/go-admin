package config

import "time"

type Server struct {
	Model        string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
