package db

import "simpleNetdisk/util/config"

const (
	// CONN_TIMEOUT TODO
	CONN_TIMEOUT string = "20s"
	// READ_TIMEOUT TODO
	READ_TIMEOUT string = "5m"
	// MYSQL_BASE TODO
	MYSQL_BASE string = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&readTimeout=%s&timeout=%s"
)

// Db TODO
type Db struct {
	Driver          string
	Host            string
	Port            int
	Database        string
	User            string
	Password        string
	Schema          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

var (
	dbConf DbConfig
)

// DbConfig TODO
type DbConfig struct {
	databases map[string]Db
}

func GetDbCfg(dbCfgKey string) (db Db, err error) {
	if _, ok := dbConf.databases[dbCfgKey]; !ok {
		dbConf.databases = make(map[string]Db)
		dbCfg := config.GlobalConfig().DB
		dbConf.databases[dbCfgKey] = Db{
			Driver:          dbCfg.Driver,
			Host:            dbCfg.Host,
			Port:            dbCfg.Port,
			Database:        dbCfg.Database,
			User:            dbCfg.User,
			Password:        dbCfg.Password,
			Schema:          dbCfg.Schema,
			MaxIdleConns:    dbCfg.MaxIdleConns,
			MaxOpenConns:    dbCfg.MaxOpenConns,
			ConnMaxLifetime: dbCfg.ConnMaxLifetime,
		}
	}
	return dbConf.databases[dbCfgKey], nil
}
