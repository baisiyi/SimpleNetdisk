package minio

import "github.com/minio/minio-go/v7"

type options struct {
	bucketKey string
}

type Option func(o *options)

type MinioHelper struct {
	client *minio.Client
	opts   *options
}

func WithBucketKey(bucketKey string) Option {
	return func(o *options) {
		o.bucketKey = bucketKey
	}
}

//NewMiniHelper 构造函数
func NewMiniHelper(client *minio.Client, opts ...Option) *MinioHelper {
	helperOts := &options{}
	for _, o := range opts {
		o(helperOts)
	}
	hp := &MinioHelper{
		client: client,
		opts:   helperOts,
	}
	return hp
}
