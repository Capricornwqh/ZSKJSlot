package Common

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

const (
	MQ_VERSION_KEY = "version" // 服務版本key
)

var NatsConn *nats.Conn

// MqInit message queue初始化
func MqInit() error {
	var err error
	NatsConn, err = nats.Connect(
		fmt.Sprintf("nats://%s:%d", Config.Nats.Host, Config.Nats.Port),
		nats.Token(Config.Nats.Token),
		nats.MaxReconnects(-1), // 不限制重連次數避免遊戲服務直接結束程式
	)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	return err
}

// RegisterService 註冊服務至message queue
func RegisterService(serviceName string, serviceVersion string, handler func([]byte) []byte) error {
	_, err := NatsConn.Subscribe(serviceName, func(m *nats.Msg) {
		respMsg := nats.NewMsg(m.Reply)
		respMsg.Data = handler(m.Data)
		respMsg.Header.Add(MQ_VERSION_KEY, serviceVersion)
		m.RespondMsg(respMsg)
	})
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	return err
}
