package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config : cacktaildb api
type Config struct {
	Env   string
	Mysql struct {
		Local struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			DB       string `yaml:"db"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"local"`
		Prod struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			DB       string `yaml:"db"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"prod"`
	} `yaml:"mysql"`
	JWT struct {
		Local struct {
			Secret     string `yaml:"secret"`
		} `yaml:"local"`
		Prod struct {
			Secret     string `yaml:"secret"`
		} `yaml:"prod"`
	} `yaml:"jwt"`

	DBCred struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       string `yaml:"db"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
	Secret string
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string, env string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// Init new YAML decode
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
