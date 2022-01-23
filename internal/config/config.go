package config

import "time"

type Config struct {
	BID          string
	LoopInterval time.Duration
}

func GetConfig() (*Config, error) {
	c := Config{}
	var err error

	c.BID, err = getEnvStr("APP_BID")
	if err != nil {
		return nil, err
	}

	loopInterval, err := getEnvIntWithDefault("APP_LOOP_INTERVAL", 5*60)
	if err != nil {
		return nil, err
	}
	c.LoopInterval = time.Duration(loopInterval) * time.Second

	return &c, nil
}

var Cfg *Config

func LoadConfig() error {
	if Cfg != nil {
		return nil
	}

	var err error
	Cfg, err = GetConfig()
	if err != nil {
		return err
	}

	return nil
}
