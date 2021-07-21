package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	vp *viper.Viper
}

var sections = make(map[string]interface{})

func NewConfig(config string) (*Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("./")
	if config != "" {
		vp.AddConfigPath(config)
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	c := &Config{vp}
	c.WatchSettingChange()
	return c, nil
}

func (c *Config) ReadConfig(k string, v interface{}) error {
	err := c.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (c *Config) ReloadAllConfig() error {
	for k, v := range sections {
		err := c.ReadConfig(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) WatchSettingChange() {
	go func() {
		c.vp.WatchConfig()
		c.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = c.ReloadAllConfig()
		})
	}()
}
