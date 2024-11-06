package conf

import env "github.com/caarlos0/env/v11"

type envdef struct {
	IP      string `env:"IP"`
	Port    string `env:"PORT"`
	Mode    string `env:"MODE"`
	RPCPort string `env:"RPCPORT"`
}

var envs envdef

func GetEnv() {
	err := env.Parse(&envs)
	if err != nil {
		panic(err)
	}
}
