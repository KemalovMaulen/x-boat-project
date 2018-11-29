package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger interface {
	Info(args ...interface{})

	Warn(args ...interface{})

	Fatal(args ...interface{})

	Debug(args ...interface{})
}

func New(ll uint8) Logger {
	return &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.Level(ll),
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
		},
	}
}
