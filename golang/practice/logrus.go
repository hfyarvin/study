//日志模块
package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/shiena/ansicolor"
)

var logExam = log.New()

func main()  {
	log.Info("hello, world")
	test1()
}

//自建logger实例
func test1()  {
	// force colors on for TextFormatter
	log.Formatter = &log.TextFormatter{
		ForceColors: true,
	}

	// then wrap the log output with it
	log.SetOutPut(ansicolor.NewAnsiColorWriter(os.Stdout))

	logExam.Out = os.Stdout

	logExam.WithFields(log.Fields{
		"animal": "walrus",
		"size": 10,
	}).Info("A group of walrus emerges from the ocean!")
}

// 输出到本地文件系统
