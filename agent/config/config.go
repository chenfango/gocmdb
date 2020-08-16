package config

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	// 保证agent上UUID不变
	UUID string
	UUIDFile string

	EndPoint string
	Token string

	LogFile string

	PID int
	PidFile string

	Heartbeat chan interface{}
	Register chan interface{}
	Log chan interface{}


}


func NewConfig(configReader *viper.Viper) (*Config, error){
	UUIDFile := configReader.GetString("uuidfile")
	fmt.Println(UUIDFile)
	if UUIDFile == ""{
		UUIDFile = "agentd.uuid"
	}

	PidFile := configReader.GetString("pidfile")
	fmt.Println(PidFile)
	if PidFile == ""{
		PidFile = "agentd.pid"
	}

	LogFile := configReader.GetString("logfile")
	fmt.Println(LogFile)
	if LogFile == ""{
		LogFile = "logs/agent.log"
	}

	EndPoint := configReader.GetString("endpoint")
	if EndPoint == ""{
		EndPoint = "http://localhost:8888/v2/api"
	}


	Token := configReader.GetString("token")
	if Token == ""{
		Token = "abc1234567"
	}


	UUID := ""

	if cxt, err := ioutil.ReadFile(UUIDFile);err==nil{
		UUID = string(cxt)
	}else if os.IsNotExist(err){
		UUID = strings.ReplaceAll(uuid.New().String(), "-", "")
		ioutil.WriteFile(UUIDFile, []byte(UUID), os.ModePerm)
	}else{
		return nil, err
	}

	PID := os.Getpid()

	ioutil.WriteFile(PidFile, []byte(strconv.Itoa(PID)), os.ModePerm)

	return &Config{
		EndPoint: EndPoint,
		UUID: UUID,
		UUIDFile: UUIDFile,
		LogFile: LogFile,
		PID: PID,
		PidFile: PidFile,
		Token: Token,
		Heartbeat: make(chan interface{}, 64),
		Register: make(chan interface{}, 64),
		Log: make(chan interface{}, 10240),
	},nil


}


