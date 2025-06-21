package Common

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type ServerConfig struct {
	ServerID        int
	CmdPath         string
	TimeZone        string
	ServerPort      int
	Release         bool
	ReadBufferSize  int
	WriteBufferSize int
	ClientMax       int
	ClientIdleTime  int

	Nats NatsConfig
}

type NatsConfig struct {
	Host  string
	Port  int
	Token string
}

const ConfigPath = "../config.toml"

var Config ServerConfig
var ServerLoc *time.Location

// ConfigInit 設定檔初始化
func ConfigInit() error {
	// 設定檔
	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	err = toml.Unmarshal(data, &Config)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	fmt.Printf("%#v\n", Config)

	// 時間物件
	ServerLoc, err = time.LoadLocation(Config.TimeZone)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
