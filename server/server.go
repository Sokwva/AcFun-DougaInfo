package server

import (
	"sokwva/acfun/dougaInfo/conf"
	"sokwva/acfun/dougaInfo/server/http"
	"sokwva/acfun/dougaInfo/server/rpc"
)

func Start() {
	switch conf.Conf.Server.Mode {
	case "http":
		http.Start()
	case "all":
		http.Start()
		rpc.Start()
	case "rpc":
		rpc.Start()
	}
}
