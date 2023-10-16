package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := setOutPutFile()
	if LogrusObj != nil {
		LogrusObj.Out = src
		return
	}
	//实例化
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel) // 设置日志级别
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2022-10-15 14:30:55",
	})
	LogrusObj = logger

}

func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil { //获取工作目录
		logFilePath = dir + "/logs"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			log.Print(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2023-10-15") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(fileName, 0777); err != nil {
			log.Print(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return src, nil
}
