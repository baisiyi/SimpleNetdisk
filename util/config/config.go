package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
	"strings"
	"sync/atomic"
)

var (
	DefaultConfigPath = "../conf/grpc.yaml"
)

var globalConfig atomic.Value

func newConfig() *Config {
	return &Config{}
}
func init() {
	globalConfig.Store(newConfig())
}

func GlobalConfig() *Config {
	return globalConfig.Load().(*Config)
}

func SetGlobalConfig(cfg *Config) {
	globalConfig.Store(cfg)
}

func LoadConfig(configPath string) (*Config, error) {
	cfg, err := parseConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func parseConfigFromFile(configPath string) (*Config, error) {
	var cfg = newConfig()
	config, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(config, cfg)
	if err != nil {
		return nil, err
	}
	SetGlobalConfig(cfg)
	return nil, nil
}

func GetServiceNetWork() string {
	return GlobalConfig().Service.Network
}

func GetServiceAddress() string {
	var s strings.Builder
	if GlobalConfig().Service.IP != "0.0.0.0" {
		s.WriteString(GlobalConfig().Service.IP)
	}
	s.WriteString(":")
	s.WriteString(strconv.Itoa(GlobalConfig().Service.Port))
	return s.String()
}
