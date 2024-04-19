package bxylog

import (
	"fmt"
	"github.com/bai-xiao-yuan/bxylog/conf"
	"runtime"
	"sync"
	"time"
)

const (
	white  = "\033[1;37m"
	Blue   = "\033[34m"
	Yellow = "\033[0;33m"
	Red    = "\033[0;31m"
	Green  = "\033[0;32m"
	Reset  = "\033[0m"
)

const layout = "2006-01-02 15:04:05.000000"

var levelMap = map[conf.LogLevel]string{
	conf.LInfo:  "[Info]: ",
	conf.LDebug: "[Debug]: ",
	conf.LWarn:  "[Warn]: ",
	conf.LError: "[Error]: ",
	conf.LPanic: "[Panic]: ",
}

var colorMap = map[conf.LogLevel]string{
	conf.LInfo:  white,
	conf.LDebug: Blue,
	conf.LWarn:  Yellow,
	conf.LError: Red,
	conf.LPanic: Green,
}

var logger = NewDefaultLog()

type BxyLog interface {
	Info(format string, v ...any)
	Debug(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
	Panic(format string, v ...any)
	print(calldepth int, level conf.LogLevel, format string, v ...any)
}

type Log struct {
	mu     sync.Mutex
	config *conf.Config
	out    *multiWriter
	buf    []byte
	w      *watch
}

func NewLog(c *conf.Config) BxyLog {
	log := &Log{
		config: c,
	}
	return log.init()
}

func NewDefaultLog() BxyLog {
	c := &conf.Config{
		Level:     conf.LInfo,
		OutTarget: conf.Std,
	}
	log := &Log{
		config: c,
	}
	return log.init()
}

func (l *Log) Info(format string, v ...any) {
	if !l.checkLevel(conf.LInfo) {
		return
	}
	format = l.formatString(format, v...)
	l.outPut(2, format, conf.LInfo)
}

func (l *Log) Debug(format string, v ...any) {
	if !l.checkLevel(conf.LDebug) {
		return
	}
	format = l.formatString(format, v...)
	l.outPut(2, format, conf.LDebug)
}

func (l *Log) Warn(format string, v ...any) {
	if !l.checkLevel(conf.LWarn) {
		return
	}
	format = l.formatString(format, v...)
	l.outPut(2, format, conf.LWarn)
}

func (l *Log) Error(format string, v ...any) {
	if !l.checkLevel(conf.LError) {
		return
	}
	format = l.formatString(format, v...)
	l.outPut(2, format, conf.LError)
}

func (l *Log) Panic(format string, v ...any) {
	if !l.checkLevel(conf.LPanic) {
		return
	}
	format = l.formatString(format, v...)
	l.outPut(2, format, conf.LPanic)
}

func (l *Log) print(calldepth int, level conf.LogLevel, format string, v ...any) {
	format = l.formatString(format, v...)
	l.outPut(calldepth, format, level)
}

func (l *Log) SetLevel(level conf.LogLevel) {
	l.config.Level = level
}

func (l *Log) SetTimeFlag(flag bool) {
	l.config.TimeFlag = flag
}

func (l *Log) SetColor(flag bool) {
	l.config.Color = flag
}

func (l *Log) init() *Log {
	l.w = newWatch(l)
	return l
}

func (l *Log) SetPrefix(prefix string) {
	l.config.Prefix = prefix
}

func (l *Log) outPut(calldepth int, s string, level conf.LogLevel) error {
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = l.buf[:0]
	l.setHande(&l.buf, file, line, level)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, err := l.out.Write(l.buf)
	return err
}

func (l *Log) setHande(buf *[]byte, file string, line int, level conf.LogLevel) {
	if l.config.Color {
		*buf = append(*buf, colorMap[level]...)
	}
	*buf = append(*buf, levelMap[level]...)
	*buf = append(*buf, l.config.Prefix...)
	if l.config.TimeFlag {
		now := time.Now()
		formattedTime := now.Format(layout)
		*buf = append(*buf, formattedTime...)
		*buf = append(*buf, ": "...)
	}
	*buf = append(*buf, file...)
	*buf = append(*buf, ':')
	itoa(buf, line, -1)
	*buf = append(*buf, ": "...)
}

func (l *Log) checkLevel(level conf.LogLevel) bool {
	if level < l.config.Level {
		return false
	}
	return true
}

func (l *Log) formatString(format string, v ...any) string {
	if len(format) == 0 {
		format = fmt.Sprintln(format)
	} else {
		format = fmt.Sprintf(format, v...)
	}
	return format
}

func (l *Log) Close() {

}

func Info(format string, v ...any) {
	logger.print(3, conf.LInfo, format, v...)
}

func Debug(format string, v ...any) {
	logger.print(3, conf.LDebug, format, v...)
}

func Warn(format string, v ...any) {
	logger.print(3, conf.LWarn, format, v...)
}

func Error(format string, v ...any) {
	logger.print(3, conf.LError, format, v...)
}

func Panic(format string, v ...any) {
	logger.print(3, conf.LPanic, format, v...)
}

func itoa(buf *[]byte, i int, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}
