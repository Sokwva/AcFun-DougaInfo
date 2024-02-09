package redis

import (
	"sokwva/acfun/dougaInfo/conf"
	"time"

	"github.com/go-redis/redis"
)

type RedisDrv struct {
	handle *redis.Client
}

func (me *RedisDrv) Connect(svrType string, host string, port string, username string, password string) {
	me.handle = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.Cache.Host + ":" + conf.Conf.Cache.Port,
		Password: conf.Conf.Cache.Password,
	})
}

func (me *RedisDrv) Ping() error {
	_, err := me.handle.Ping().Result()
	return err
}

func (me *RedisDrv) Get(name string) (interface{}, error) {
	return me.handle.Get(conf.Conf.Cache.Prefix + name).Result()
}

func (me *RedisDrv) Set(name string, value interface{}, expireTime time.Duration) {
	if expireTime == 0 {
		me.handle.Set(conf.Conf.Cache.Prefix+name, value, time.Duration(conf.Conf.Cache.ExpireDuration)*time.Second)
		return
	}
	me.handle.Set(conf.Conf.Cache.Prefix+name, value, expireTime*time.Second)
}
