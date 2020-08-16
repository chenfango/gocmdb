package main

import (
	"github.com/chenfanlinux/gocmdb/agent/config"
	"github.com/chenfanlinux/gocmdb/agent/ens"
	"github.com/chenfanlinux/gocmdb/agent/plugins"
	_ "github.com/chenfanlinux/gocmdb/agent/plugins/init"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func main(){
	logrus.SetLevel(logrus.DebugLevel)
	configReader := viper.New()
	configReader.SetConfigName("agent")
	configReader.SetConfigType("yaml")
	configReader.AddConfigPath("etc/")
	err := configReader.ReadInConfig()
	if err!=nil{
		logrus.Error("读取配置出错:", err)
		os.Exit(-1)
	}

	gconf, err := config.NewConfig(configReader)


	defer func() {
		os.Remove(gconf.PidFile)
	}()

	logrus.SetLevel(logrus.DebugLevel)
	log, err := os.OpenFile(gconf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err !=nil{
		logrus.Error("打开日志文件出错:", err)
		os.Exit(-1)
	}
	//logrus.SetOutput(log)

	defer func() {
		log.Close()
	}()

	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetOutput(log)
	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent启动")


	plugins.DefaultManager.Init(gconf)

	ens.NewENS(gconf).Start()


	plugins.DefaultManager.Start()


	// 让程序一直运行起来，不会退出，工作例程会随着主例程退出而退出
	go func() {
		for now := range  time.Tick(time.Second){
			logrus.Info(now)
		}
	}()

	// 让程序等待信号
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch

	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent退出")




}

