package engine

import "github.com/denkhaus/yamlconfig"

type Config struct {
	*yamlconfig.Config
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (c *Config) GetMongoConfig() (map[string]interface{}, error) {
	conf := make(map[string]interface{})
	conf["host"] = c.GetString("mongo:host")
	return conf, nil
}

////////////////////////////////////////////////////////////////////////////////
func (c *Config) Init(cnfName string) error {
	c.Config = yamlconfig.NewConfig(cnfName)

	if err := c.Load(func(config *yamlconfig.Config) {
		config.SetDefault("mongo:host", "127.0.0.1")
	}, "", false); err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
func NewConfig(cnfName string) (*Config, error) {
	cfig := &Config{}
	if err := cfig.Init(cnfName); err != nil {
		return nil, err
	}
	return cfig, nil
}
