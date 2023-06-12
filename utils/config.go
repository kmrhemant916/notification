package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
    Rabbitmq struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
        Exchange struct {
            Name string `yaml:"name"`
            Kind string `yaml:"kind"`
        } `yaml:"exchange"`
	} `yaml:"rabbitmq"`
}

func (c *Config) ReadConf(f string) (*Config, error) {
    buf, err := ioutil.ReadFile(f)
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(buf, c)
    if err != nil {
        return nil, fmt.Errorf("in file %q: %w", f, err)
    }
    return c, err
}