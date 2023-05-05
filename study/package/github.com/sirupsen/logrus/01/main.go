package main

import (
	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	hookConfig := logredis.HookConfig{
		Host:     "localhost",
		Key:      "key001",
		Format:   "v0",
		App:      "app-sxy",
		Hostname: "redis",
		TTL:      3600,
		DB:       0,
		Port:     6379,
	}
	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error :%q", err)
	}
}
func main() {
	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	logrus.Info("just some info logging...")
	logrus.WithFields(logrus.Fields{
		"animal": "cat",
		"foo":    "bar",
		"name":   "sss",
	}).Info("additional fields are being logged as well")
	logrus.SetOutput(file)
	logrus.Info("this will only be sent to redis")
}
