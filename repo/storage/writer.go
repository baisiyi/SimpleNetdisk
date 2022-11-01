package storage

import (
	"simpleNetdisk/util/db"
	"simpleNetdisk/util/minio"
)

type writerImpl struct {
	hp        *minio.MinioHelper
	bucketKey string
}

func newWriterImpl() *writerImpl {
	return &writerImpl{}
}

func (w *writerImpl) getMiniHelper() (hp *minio.MinioHelper, err error) {
	if w.hp == nil {
		client, err := minio.NewMinioWithLocalCfg()
		if err != nil {
			return
		}
		w.hp = minio.NewMiniHelper(client, minio.WithBucketKey(w.bucketKey))
	}
	return w.hp, nil
}

func (w *writerImpl) getDbHelper() (hp *db.DbHelper, err error) {
	if w.hp == nil {
		dbInstance, err := db.NewDbWithLocalCfg("")
		if err != nil {
			return
		}
		_ = db.NewDBHelper(dbInstance.DB, db.WithTableName(""))
	}
	return nil, nil
}
