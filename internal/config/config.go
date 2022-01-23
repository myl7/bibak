package config

type Config struct {
	BID string
}

func GetConfig() (*Config, error) {
	c := Config{}
	var err error

	c.BID, err = getEnvStr("APP_BID")
	if err != nil {
		return nil, err
	}

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
