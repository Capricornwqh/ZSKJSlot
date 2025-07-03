package utils

import (
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
)

var GeoDB *geoip2.Reader

// 初始化Geo
func SetupGeo() {
	// 打开GeoIP2数据库
	db, err := geoip2.Open(Conf.GeoDB)
	if err != nil {
		logrus.Fatal(err)
	}
	GeoDB = db
}

// 关闭Geo
func CloseGeo() error {
	if GeoDB != nil {
		err := GeoDB.Close()
		if err != nil {
			return err
		}
	}
	GeoDB = nil
	return nil
}

// 获取IP所在国家
func GetIPCountry(addr, locale string) string {
	ip := net.ParseIP(addr)
	record, err := GeoDB.Country(ip)
	if err != nil {
		return ""
	}

	return record.Country.Names[locale]
}

// 获取IP所在城市
func GetIPCity(addr, locale string) string {
	ip := net.ParseIP(addr)
	record, err := GeoDB.City(ip)
	if err != nil {
		return ""
	}

	return record.City.Names[locale]
}
