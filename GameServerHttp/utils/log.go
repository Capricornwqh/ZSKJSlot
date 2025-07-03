package utils

import (
	"github.com/sirupsen/logrus"
)

// 日志配置
func SetupLogging() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.00000000",
		// FullTimestamp:    true,
		// DisableColors:    false,
		DisableTimestamp: false,
		PrettyPrint:      true,
		// QuoteEmptyFields: true,
		// CallerPrettyfier: callerPrettyfier,
	})

	if len(Conf.Logging) == 1 {
		level, err := logrus.ParseLevel(Conf.Logging[0].Level)
		if err != nil {
			logrus.Fatalf("Unrecognised logging level %s: %q", Conf.Logging[0].Level, err)
		}
		logrus.SetLevel(level)
	}
}
