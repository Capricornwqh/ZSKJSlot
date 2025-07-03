package service

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

type GameService struct {
}

// 检查所有属性是否已初始化
func (s *GameService) CheckInitialization() {
	v := reflect.ValueOf(s).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsNil() {
			logrus.Fatalf("%s field %s is not initialized", t.Name(), t.Field(i).Name)
		}
	}
}
