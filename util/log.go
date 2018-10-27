package util

import (
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"sync"
	"tyrannosaurs/config"
	"fmt"
	"flag"
)

var Log *logrus.Logger
var once sync.Once

func init() {
	flag.Parse()
	EnvConfig, err := config.Env.GetConfig(*config.E)
	if err != nil {
		panic(err)
	}
	once.Do(func() {
		pathMap := lfshook.PathMap{
			logrus.InfoLevel:  fmt.Sprintf("%s/info.log", EnvConfig.LogPath),
			logrus.ErrorLevel: fmt.Sprintf("%s/error.log", EnvConfig.LogPath),
		}
		Log = logrus.New()
		Log.Hooks.Add(lfshook.NewHook(
			pathMap,
			&logrus.JSONFormatter{},
		))
	})
}


