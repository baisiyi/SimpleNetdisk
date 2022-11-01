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
	DB struct {
		Driver          string `yaml:"driver"`
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		Database        string `yaml:"database"`
		User            string `yaml:"user"`
		Password        string `yaml:"password"`
		Schema          string `yaml:"schema"`
		MaxIdleConns    int    `yaml:"maxIdleConns"`
		MaxOpenConns    int    `yaml:"maxOpenConns"`
		ConnMaxLifetime int    `yaml:"connMaxLifetime"`
	}
}
