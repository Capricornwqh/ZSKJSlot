package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PostgresClub = iota
	PostgresSpin
)

var PostgresDB [2]*gorm.DB

// 初始化数据库
func SetupPostgreSQL() {
	// pConfig := config.Conf.PostgreSQL
	clubDB, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s connect_timeout=10 TimeZone=Asia/Shanghai",
			Conf.PostgreSQL.Game.Host, Conf.PostgreSQL.Game.UserName, Conf.PostgreSQL.Game.Password, Conf.PostgreSQL.Game.Database,
			Conf.PostgreSQL.Game.Port, Conf.PostgreSQL.Game.SSLMode)),
		&gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	if dbase, err := clubDB.DB(); err != nil {
		logrus.Fatal(err)
	} else {
		if err := dbase.Ping(); err != nil {
			logrus.Fatal(err)
		}
		// dbase.SetMaxIdleConns(20)
		// dbase.SetMaxOpenConns(500)
		// dbase.SetConnMaxLifetime(90 * time.Second)
		PostgresDB[PostgresClub] = clubDB
		logrus.Infof("PostgreSQL Club Server Open: username: %s, address: %s:%d",
			Conf.PostgreSQL.Game.UserName, Conf.PostgreSQL.Game.Host, Conf.PostgreSQL.Game.Port)
	}

	spinDB, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s connect_timeout=10 TimeZone=Asia/Shanghai",
			Conf.PostgreSQL.Log.Host, Conf.PostgreSQL.Log.UserName, Conf.PostgreSQL.Log.Password, Conf.PostgreSQL.Log.Database,
			Conf.PostgreSQL.Log.Port, Conf.PostgreSQL.Log.SSLMode)),
		&gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	if dbase, err := spinDB.DB(); err != nil {
		logrus.Fatal(err)
	} else {
		if err := dbase.Ping(); err != nil {
			logrus.Fatal(err)
		}

		PostgresDB[PostgresSpin] = spinDB
		logrus.Infof("PostgreSQL Spin Server Open: username: %s, address: %s:%d",
			Conf.PostgreSQL.Log.UserName, Conf.PostgreSQL.Log.Host, Conf.PostgreSQL.Log.Port)
	}
}

func ClosePostgreSQL() error {
	for _, db := range PostgresDB {
		if db == nil {
			continue
		}
		if dbase, err := db.DB(); err != nil {
			logrus.Error(err)
		} else {
			if err := dbase.Close(); err != nil {
				logrus.Error(err)
			}
		}
		db = nil
	}

	return nil
}
