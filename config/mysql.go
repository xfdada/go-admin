package config

import "time"

type Mysql struct {
	ParseTime       bool
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
	Username        string
	Password        string
	Host            string
	DBName          string
	TablePrefix     string
	Charset         string
}
