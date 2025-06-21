package main

import (
	"Force/GameServer/Common"
	"Force/GameServer/Slot013FatToad"
	"runtime"
	"strconv"
)

// RELEASE_MODE Release 模式
var RELEASE_MODE string

func main() {
	Slot013FatToad.RELEASE_MODE, _ = strconv.ParseBool(RELEASE_MODE)

	err := Common.ConfigInit()
	checkErr(err)

	err = Common.MqInit()
	checkErr(err)

	err = Common.RegisterService("FatToad", Slot013FatToad.PROB_VERSION, Slot013FatToad.MessageHandler)
	checkErr(err)

	runtime.Goexit()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
