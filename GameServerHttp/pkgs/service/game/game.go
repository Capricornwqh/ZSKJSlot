package service_game

import (
	"reflect"

	"github.com/gin-gonic/gin"
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

// 游戏列表
func (s *GameService) GameList(ctx *gin.Context) {

}
