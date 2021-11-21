package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DatasourceCardanoDbSync struct {
	Host     string `yaml:"host" default:"localhost"`
	Port     int    `yaml:"port" default:"5432"`
	Username string `yaml:"username" default:"postgres"`
	Password string `yaml:"password" default:"changeme"`
	DbName   string `yaml:"dbname" default:"cexplorer"`
	SslMode  string `yaml:"sslmode" default:"prefer"`
}

type Datasource struct {
	CardanoDbSync DatasourceCardanoDbSync `yaml:"cardanodbsync"`
}

type Api struct {
	Address string `yaml:"address" default:"0.0.0.0"`
	Port    int    `yaml:"port" default:"8080"`
}

type Config struct {
	Api        Api        `yaml:"api"`
	Datasource Datasource `yaml:"datasource"`
}

func New(configFile string) (*Config, error) {
	c := &Config{}
	if configFile != "" {
		buf, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(buf, c)
		if err != nil {
			return nil, err
		}
	}
	err := envconfig.Process("", c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
