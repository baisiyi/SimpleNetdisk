package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var newLogger logger.Interface

// 管理连接，确保每个db单例, 兼容多种不同类型DB
var dbManager = new(sync.Map)

// OneInstance DB实例
type OneInstance struct {
	DB   *gorm.DB
	once *sync.Once
}

func init() {
	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
}

//NewDbWithLocalCfg 按照指定本地配置新建DB实例
func NewDbWithLocalCfg(dbCfgKey string) (dbInstance *OneInstance, err error) {
	if val, ok := dbManager.Load(dbCfgKey); ok {
		dbInstance = val.(*OneInstance)
	} else {
		dbInstance = new(OneInstance)
		dbInstance.once = new(sync.Once)
		val1, _ := dbManager.LoadOrStore(dbCfgKey, dbInstance)
		dbInstance = val1.(*OneInstance)
	}
	dbInstance.once.Do(func() {
		var db *gorm.DB
		db, err = getConn(dbCfgKey)
		if err != nil {
			log.Fatal(err)
			return
		}
		dbInstance.DB = db
	})
	return
}

// getConn 获取数据库链接
func getConn(dbCfgKey string) (db *gorm.DB, err error) {
	cfg, err := GetDbCfg(dbCfgKey)
	if err != nil {
		return
	}
	switch cfg.Driver {
	case "mysql":
		dsn := fmt.Sprintf(MYSQL_BASE, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, READ_TIMEOUT, CONN_TIMEOUT)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
		break
	}

	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	return
}
