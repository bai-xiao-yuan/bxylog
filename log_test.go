package bxylog

import (
	"testing"

	"github.com/bai-xiao-yuan/bxylog/conf"
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
		Level:         conf.LInfo,
		OutTarget:     conf.File,
		FileName:      "log/t1.log",
		Color:         true,
		Prefix:        "TestLogger: ",
		TimeFlag:      true,
		FileSliceType: conf.FileSize,
		SliceProp:     1024,
	})
	for i := 0; i < 10000; i++ {
		l.Info("测试等级Info")
		l.Debug("测试等级Debug")
	}
}
