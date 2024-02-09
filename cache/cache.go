package cache

import (
	"errors"
	"sokwva/acfun/dougaInfo/cache/empty"
	"sokwva/acfun/dougaInfo/cache/redis"
	"sokwva/acfun/dougaInfo/conf"
	"time"
)

type CacheDriver interface {
	Connect(svrType string, host string, port string, username string, password string)
	Ping() error
	Get(target string) (interface{}, error)
	Set(name string, value interface{}, expireDuration time.Duration)
}

var Handle CacheDriver

func Start() error {
	if !conf.Conf.Cache.Enable {
		Handle = &empty.EmptyDrv{}
		return errors.New("empty")
	}
	switch conf.Conf.Cache.ServerType {
	case "redis":
		Handle = &redis.RedisDrv{}
		Handle.Connect("", conf.Conf.Cache.Host, conf.Conf.Cache.Port, "", "")
		return Handle.Ping()
	}
	return errors.New("did not select a cache server")
}
