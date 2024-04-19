package bxylog

import (
	"bxylog/conf"
	"bxylog/util"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const FileFix = "2006-01-02_15_04_05"
const dateFormat = "2006-01-02"

type watch struct {
	l    *Log
	file *os.File
	time time.Time
}

func newWatch(l *Log) *watch {
	w := watch{
		l: l,
	}
	return w.Init()
}

func (w *watch) Init() *watch {
	w.newWriter()
	switch w.l.config.FileSliceType {
	case conf.FileSize:
		go w.watchSize()
	case conf.Day:
		go w.watchTime()
	case conf.TimeDate:
		go w.watchDay()
	}
	return w
}

func (w *watch) newFile() {
	t := time.Now()
	fileName := w.l.config.FileName
	if w.l.config.FileSliceType == conf.Day {
		parse, err := time.Parse(dateFormat, t.Format(dateFormat))
		if err != nil {
			log.Fatalf("time Parse error: %v", err)
		}
		w.time = parse
	} else if w.l.config.FileSliceType == conf.TimeDate {
		w.time = t
	}

	if w.l.config.FileSliceType != conf.None {
		files := strings.Split(w.l.config.FileName, ".")
		fileName = fmt.Sprintf("%s_%s.%s", files[0], t.Format(FileFix), files[1])
	}
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	w.file = f
}

func (w *watch) newWriter() {
	if w.l.config.OutTarget == conf.File {
		dirName := filepath.Dir(w.l.config.FileName)
		util.MakeDir(dirName)
		w.newFile()
		w.l.out = &multiWriter{
			writers: []io.Writer{
				os.Stdout, // 控制台输出
				w.file,    // 文件输出
			},
		}
	} else if w.l.config.OutTarget == conf.Std {
		w.l.out = &multiWriter{
			writers: []io.Writer{
				os.Stdout, // 控制台输出
			},
		}
	}
}

func (w *watch) changeFile() {
	w.Close()
	w.newFile()
	w.l.out.ChangeFile(w.file)
}

func (w *watch) watchSize() {
	if w.l.config.SliceProp == 0 {
		return
	}
	for true {
		info, err := w.file.Stat()
		if err != nil {
			log.Fatalf("watchSize errpr: %v", err)
		}
		if info.Size() >= w.l.config.SliceProp*1024 {
			w.changeFile()
		}
	}

}

func (w *watch) watchTime() {
	if w.l.config.SliceProp == 0 {
		return
	}
	for true {
		if time.Now().Unix()-w.time.Unix() >= w.l.config.SliceProp*3600 {
			w.changeFile()
		}
	}

}

func (w *watch) watchDay() {
	if w.l.config.SliceProp == 0 {
		return
	}
	for true {
		if time.Now().Unix()-w.time.Unix() >= w.l.config.SliceProp*3600*24 {
			w.changeFile()
		}
	}
}

func (w *watch) Close() {
	if w.file != nil {
		err := w.file.Close()
		if err != nil {
			log.Fatalf("close log file error: %v", err)
		}
	}
}
