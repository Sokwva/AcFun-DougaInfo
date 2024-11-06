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
	RPCPort string
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
		GetEnv()
		Conf = MainConf{
			Server: SvrConf{
				Address: envs.IP,
				Port:    envs.Port,
				RPCPort: envs.RPCPort,
				Mode:    envs.Mode,
			},
			Common: CommenConf{
				LogLevel: "info",
			},
			Cache: CacheConf{
				Enable:         false,
				ServerType:     "",
				Host:           "",
				Port:           "",
				UserName:       "",
				Prefix:         "",
				ExpireDuration: 1,
			},
		}
	}
}
