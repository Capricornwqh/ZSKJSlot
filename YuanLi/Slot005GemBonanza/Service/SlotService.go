package main

import (
	"Force/GameServer/Common"
	"Force/GameServer/Slot005GemBonanza"
	"runtime"
	"strconv"
)

// RELEASE_MODE Release 模式
var RELEASE_MODE string

func main() {
	Slot005GemBonanza.RELEASE_MODE, _ = strconv.ParseBool(RELEASE_MODE)

	err := Common.ConfigInit()
	checkErr(err)

	err = Common.MqInit()
	checkErr(err)

	err = Common.RegisterService("GemBonanza", Slot005GemBonanza.PROB_VERSION, Slot005GemBonanza.MessageHandler)
	checkErr(err)

	runtime.Goexit()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
