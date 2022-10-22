package config

type Config struct {
	Service struct {
		IP      string `yaml:"ip"`
		Port    int    `yaml:"port"`
		Network string `yaml:"network"`
	} `yaml:"service"`
}
