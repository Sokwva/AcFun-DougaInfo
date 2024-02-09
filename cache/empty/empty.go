package empty

import (
	"errors"
	"time"
)

type EmptyDrv struct {
	Handle interface{}
}

func (me *EmptyDrv) Connect(svrType string, host string, port string, username string, password string) {
}

func (me *EmptyDrv) Ping() error {
	return errors.New("empty")
}

func (me *EmptyDrv) Get(name string) (interface{}, error) {
	return "", errors.New("empty")
}

func (me *EmptyDrv) Set(name string, value interface{}, expireTime time.Duration) {

}
