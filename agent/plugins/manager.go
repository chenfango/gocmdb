package plugins

import (
	"github.com/chenfanlinux/gocmdb/agent/config"
	"github.com/sirupsen/logrus"
	"time"
)

type Manager struct {
	Cycles map[string]CyclePlugin
}

func NewManager() *Manager{
	return &Manager{
		Cycles:make(map[string]CyclePlugin),
	}
}

func (m *Manager) RegisterCycle(p CyclePlugin)  {
	m.Cycles[p.Name()] = p
	logrus.WithFields(logrus.Fields{
		"Name" : p.Name(),
	}).Info("插件注册")
}

func (m *Manager) Init(conf *config.Config){
	for name, plugin := range m.Cycles {
		plugin.Init(conf)
		logrus.WithFields(logrus.Fields{
			"Name": name,
		}).Info("初始化插件")
	}
}


func (m *Manager) Start(){
	go m.StartCycle()
}

func (m *Manager) StartCycle() {
	// 每秒遍历循环插件，当下一次执行时间<当前时间, 该执行了
	for now := range time.Tick(time.Second) {
		for name, plugin := range m.Cycles {
			if now.After(plugin.NextTime()) {
				if evt, err := plugin.Call(); err == nil {
					logrus.WithFields(logrus.Fields{
						"Name" : name,
						"Result" : evt,
					}).Debug("插件执行")
					// 结果写入管道
					plugin.Pipeline() <- evt

				} else {
					logrus.WithFields(logrus.Fields{
						"Name" : name,
						"error" : err,
					}).Debug("插件执行失败")
				}
			}
		}
	}
}

var DefaultManager = NewManager()

