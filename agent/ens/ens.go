package ens

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/chenfanlinux/gocmdb/agent/config"

	"github.com/imroc/req"

)

type ENS struct {
	conf *config.Config

}

func NewENS(conf *config.Config) *ENS {
	return &ENS{conf: conf}
}



func (s *ENS) Start() {
	headers := req.Header{"Token": s.conf.Token}

	logrus.Info("ENS 开始运行")
	//req.Debug = true
	go func() {
		endpoint := fmt.Sprintf("%s/heartbeat/%s/", s.conf.EndPoint, s.conf.UUID)
		for evt := range s.conf.Heartbeat {
			response, err := req.New().Post(endpoint, req.BodyJSON(evt), headers)
			if err == nil{
				result := map[string]interface{}{}
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"result": result,
				}).Debug("上传心跳信息成功")
			}else{
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("上传心跳信息失败")
			}
			logrus.Debug("", evt)
		}
	}()



	go func() {
		endpoint := fmt.Sprintf("%s/register/%s/", s.conf.EndPoint, s.conf.UUID)
		for evt := range s.conf.Register {
			response, err := req.New().Post(endpoint, req.BodyJSON(evt), headers)
			if err == nil{
				result := map[string]interface{}{}
				response.ToJSON(&result)

				logrus.WithFields(logrus.Fields{
					"result": result,
				}).Debug("注册成功")
			}else{
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("注册失败")
			}
			logrus.Debug("", evt)
		}
	}()




	go func() {
		endpoint := fmt.Sprintf("%s/log/%s/", s.conf.EndPoint, s.conf.UUID)
		for evt := range s.conf.Log {
			response, err := req.New().Post(endpoint, req.BodyJSON(evt), headers)
			if err == nil {
				result := map[string]interface{}{}
				response.ToJSON(&result)
				logrus.WithFields(logrus.Fields{
					"result" : result,
				}).Debug("日志上传成功")
			} else {
				logrus.WithFields(logrus.Fields{
					"error" : err,
				}).Error("日志上传失败")
			}
		}
	}()



}