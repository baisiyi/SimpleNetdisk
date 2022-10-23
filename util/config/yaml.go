package config

type Config struct {
	Service struct {
		IP      string `yaml:"ip"`
		Port    int    `yaml:"port"`
		Network string `yaml:"network"`
	} `yaml:"service"`
	Minio struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"accessKeyID"`
		SecretAccessKey string `yaml:"secretAccessKey"`
	} `yaml:"minio"`
}
