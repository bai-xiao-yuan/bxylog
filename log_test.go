package bxylog

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var a = sync.WaitGroup{}

func TestLog(t *testing.T) {
	// 输出彩色文本
	//fmt.Println(white, "This is red text")
	//fmt.Println(Blue, "This is red text", "This is red text")
	//fmt.Println(Yellow + "This is red text")
	//fmt.Println(Red + "This is red text")
	//fmt.Println(Green + "This is red text")
	//log.Printf("测试%v", "sdasda")
	for true {
		Info("Info")
		Debug("Debug")
		Warn("Warn")
		Error("Error")
		Panic("Panic")
	}

}

func Tes(i int) {
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println(i, "1123123")
	a.Done()
}
