package databasenani

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func SetUpDB(config *gorm.Config, dsn string, maxIdleConns, maxOpenConns int, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}
	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		}
	}
	if db, err = gorm.Open(mysql.Open(dsn), config); err != nil {
		print("opens database failed: %s", err.Error())
		return
	}
	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
	} else {
		print(err)
	}

	if err = db.AutoMigrate(models...); nil != err {
		print("auto migrate tables failed: %s", err.Error())
	}
	return
}

func CloseDB() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		print("Disconnect from database failed: %s", err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
