package logger

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger     *logrus.Logger
	loggerOnce sync.Once
)

func initLogger() *logrus.Logger {
	executablePath, err := os.Executable()
	if err != nil {
		panic("can not get the exectable path: " + err.Error())
	}

	executableDir := filepath.Dir(executablePath)

	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("can not create the log dir: " + err.Error())
	}

	logger := logrus.New()

	logger.Out = &lumberjack.Logger{
		Filename:   filepath.Join(executableDir, logDir, "simtrans.log"),
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}

	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func GetLogger() *logrus.Logger {
	loggerOnce.Do(
		func() {
			logger = initLogger()
		},
	)
	return logger
}
