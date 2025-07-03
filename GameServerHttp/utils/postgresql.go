package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgreSQLDB *gorm.DB

// 初始化数据库
func SetupPostgreSQL() {
	// pConfig := config.Conf.PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s connect_timeout=10 TimeZone=Asia/Shanghai", Conf.PostgreSQL.Host,
		Conf.PostgreSQL.UserName, Conf.PostgreSQL.Password, Conf.PostgreSQL.Database, Conf.PostgreSQL.Port, Conf.PostgreSQL.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	if dbase, err := db.DB(); err != nil {
		logrus.Fatal(err)
	} else {
		if err := dbase.Ping(); err != nil {
			logrus.Fatal(err)
		}
		// dbase.SetMaxIdleConns(20)
		// dbase.SetMaxOpenConns(500)
		// dbase.SetConnMaxLifetime(90 * time.Second)
	}

	PostgreSQLDB = db
	logrus.Infof("PostgreSQL Server Open: username: %s, address: %s:%d", Conf.PostgreSQL.UserName, Conf.PostgreSQL.Host, Conf.PostgreSQL.Port)
}

func ClosePostgreSQL() error {
	if PostgreSQLDB != nil {
		sqlDB, err := PostgreSQLDB.DB()
		if err != nil {
			return err
		}

		return sqlDB.Close()
	}

	return nil
}
