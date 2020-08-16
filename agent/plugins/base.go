package plugins

import (
	"time"
	"github.com/chenfanlinux/gocmdb/agent/config"

	)

type CyclePlugin interface {
	Name() string
	Init(*config.Config)
	NextTime() time.Time
	Call() (interface{}, error)
	Pipeline()  chan interface{}

}




