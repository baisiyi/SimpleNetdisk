package config

var (
	DefaultConfigPath = "./conf/grpc.yaml"
)

type Config struct {
}

func newConfig() *Config {
	return &Config{}
}

func LoadConfig(configPath string) (*Config, error) {
	cfg, err := parseConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func parseConfigFromFile(configPath string) (*Config, error) {
	return nil, nil
}
