package util

import (
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"sync"
)

var Log *logrus.Logger
var once sync.Once

func NewLogger() *logrus.Logger {
	once.Do(func() {
		pathMap := lfshook.PathMap{
			logrus.InfoLevel:  "./info.log",
			logrus.ErrorLevel: "./error.log",
		}
		Log = logrus.New()
		Log.Hooks.Add(lfshook.NewHook(
			pathMap,
			&logrus.JSONFormatter{},
		))
	})
	return Log
}

func init() {
	Log = NewLogger()
}
