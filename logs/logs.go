package logs

import (
	"fmt"
	"github.com/TobiasYin/goi/dom"
	"strings"
	"time"
)

const (
	Info = iota
	Debug
	Warning
	Error
	None
)

var (
	logLevel   = Info
	printLevel = Info
)

func init() {

}

func SetLogLevel(level int) {
	if level < 0 {
		level = 0
	}
	if level > None {
		level = None
	}
	logLevel = level
}

func SetPrintLevel(level int) {
	if level < 0 {
		level = 0
	}
	if level > None {
		level = None
	}
	printLevel = level
}

func Errorf(format string, args ...interface{}) {
	printfWithLevel(Error, format, args...)
}

func Errorln(args ...interface{}) {
	printlnWithLevel(Error, args...)
}

func Warningf(format string, args ...interface{}) {
	printfWithLevel(Warning, format, args...)
}

func Warningln(args ...interface{}) {
	printlnWithLevel(Warning, args...)
}

func Debugf(format string, args ...interface{}) {
	printfWithLevel(Debug, format, args...)
}

func Debugln(args ...interface{}) {
	printlnWithLevel(Debug, args...)
}

func Infof(format string, args ...interface{}) {
	printfWithLevel(Info, format, args...)
}

func Infoln(args ...interface{}) {
	printlnWithLevel(Info, args...)
}

func Println(args ...interface{}) {
	printlnWithLevel(printLevel, args...)
}

func Printf(format string, args ...interface{}) {
	printfWithLevel(printLevel, format, args...)
}

func printlnWithLevel(level int, args ...interface{}) {
	var formatStr strings.Builder
	for i, _ := range args {
		if i != len(args) {
			formatStr.WriteString("%v ")
		} else {
			formatStr.WriteString("%v")
		}
	}
	println_(level, fmt.Sprintf(formatStr.String(), args...))
}

func printfWithLevel(level int, format string, args ...interface{}) {
	println_(level, fmt.Sprintf(format, args...))
}

func println_(level int, str string) {
	if level < logLevel {
		return
	}
	go func() {
		dom.Win.Get("console").Call("log", str)
	}()
}

type Clock struct {
	startTime time.Time
	endTime   time.Time
	Level     int
	Hint      string
}

func (c *Clock) Start() {
	if c.Level < logLevel {
		return
	}
	c.startTime = time.Now()
}

func (c *Clock) End() {
	if c.Level < logLevel {
		return
	}
	c.endTime = time.Now()
	printfWithLevel(c.Level, "%s, Using: %v\n", c.Hint, c.endTime.Sub(c.startTime))
}
