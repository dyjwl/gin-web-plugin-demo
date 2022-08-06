package crontab

import (
	"time"

	"github.com/dyjwl/gin-web-plugin-demo/pkg/log"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/schedule"
)

var testCron int32 = int32(2)
var (
	dailySched = schedule.NewInShanghai("Test Cron Job", testCron)
)

func Println() error {
	log.Info("worker function")
	return nil
}

func Run() {
	dailySched.Task("Test Job").
		DisLock(5 * time.Minute).
		AddFunc(Println).
		DoCron("*/1 * * * * *")
	dailySched.Start()
}
