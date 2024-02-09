package main

import (
	"log/slog"
	"os"
	"sokwva/acfun/dougaInfo/cache"
	"sokwva/acfun/dougaInfo/conf"
	"sokwva/acfun/dougaInfo/server"
)

var slogger slog.Logger

func logLevelMap(str string) slog.Level {
	var logMap map[string]slog.Level = map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}
	return logMap[str]
}

func initLogger() {
	slogger = *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevelMap("info"),
	}))
}

func main() {
	initLogger()
	conf.InitConfDriver()
	slogger.Debug("initLogger.")
	cacheInitErr := cache.Start()
	if cacheInitErr != nil {
		slogger.Warn("cacheDrv connect faild,about 15s late would be happen. error message: " + cacheInitErr.Error())
	}
	server.Start()
}
