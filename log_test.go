package bxylog

import (
	"github.com/bai-xiao-yuan/bxylog/conf"
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	Info("Info")
	Debug("Debug")
	Warn("Warn")
	Error("Error")
	Panic("Panic")
}

func TestLogger(t *testing.T) {
	l := NewLog(&conf.Config{
		Level:         conf.LDebug,
		OutTarget:     conf.File,
		FileName:      "log/t1.log",
		Color:         true,
		Prefix:        "TestLogger: ",
		TimeFlag:      true,
		FileSliceType: conf.FileSize,
		SliceProp:     1024,
	})
	l.Info("测试等级Info")
	l.Debug("测试等级Debug")
	l.Panic("测试等级Panic")
	log.Panicln()
}
