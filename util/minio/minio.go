package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"simpleNetdisk/util/config"
	"sync"
)

var (
	once   sync.Once
	client *minio.Client
)

type Config struct {
	Endpoint string
	Options  *minio.Options
}

func getMinioCfg() *Config {
	endpoint := config.GlobalConfig().Minio.Endpoint
	accessKeyID := config.GlobalConfig().Minio.AccessKeyID
	secretAccessKey := config.GlobalConfig().Minio.SecretAccessKey
	return &Config{
		Endpoint: endpoint,
		Options: &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: true,
		},
	}
}

//NewMinioWithLocalCfg 按照指定本地配置新建DB实例
func NewMinioWithLocalCfg() (client *minio.Client, err error) {
	cfg := getMinioCfg()
	if cfg == nil {
		log.Fatal(err)
		return
	}
	return NewMinioWithCfg(cfg)
}

func NewMinioWithCfg(cfg *Config) (*minio.Client, error) {
	var err error
	once.Do(func() {
		client, err = minio.New(cfg.Endpoint, cfg.Options)
		if err != nil {
			log.Fatal(err)
		}
	})
	return client, err
}
