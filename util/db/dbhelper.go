package db

import "gorm.io/gorm"

type Option func(o *options)

type options struct {
	tableName string
}

type DbHelper struct {
	client *gorm.DB
	opts   *options
}

func WithTableName(tableName string) Option {
	return func(o *options) {
		o.tableName = tableName
	}
}

func NewDBHelper(db *gorm.DB, opts ...Option) *DbHelper {
	helperOts := &options{}
	for _, o := range opts {
		o(helperOts)
	}
	hp := &DbHelper{
		client: db,
		opts:   helperOts,
	}
	return hp
}
