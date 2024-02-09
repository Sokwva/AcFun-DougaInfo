package conf

import (
	"github.com/BurntSushi/toml"
)

var Conf MainConf

type CommenConf struct {
	LogLevel string
}

type SvrConf struct {
	Address string
	Port    string
	Mode    string
}

type CacheConf struct {
	Enable         bool
	ServerType     string
	Host           string
	Port           string
	UserName       string
	Password       string
	Prefix         string
	ExpireDuration uint
}

type MainConf struct {
	Common CommenConf
	Server SvrConf
	Cache  CacheConf
}

func InitConfDriver() {
	if _, err := toml.DecodeFile("./conf.toml", &Conf); err != nil {
		panic(err)
	}
}
