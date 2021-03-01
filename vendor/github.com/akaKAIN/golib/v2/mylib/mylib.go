package mylib

import (
	"github.com/sirupsen/logrus"
	"time"
)

func ShowCurrentTime() {
	t := time.Now().Format("2006.01.02-15.04.05")
	logrus.WithField("time", t).Info("current time from v2")
}
