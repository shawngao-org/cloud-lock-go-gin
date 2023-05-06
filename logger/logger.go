package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	Black   = "\033[1;30m"
	Red     = "\033[1;31m"
	Green   = "\033[1;32m"
	Yellow  = "\033[1;33m"
	Blue    = "\033[1;34m"
	Magenta = "\033[1;35m"
	Cyan    = "\033[1;36m"
	White   = "\033[1;37m"
	Reset   = "\033[0m"
)

func LogErr(format string, a ...any) {
	logImpl(Red+"["+getNowTimeString()+"] ERROR: "+format+"\n"+Reset, a...)
}

func LogWarn(format string, a ...any) {
	logImpl(Yellow+"["+getNowTimeString()+"] WARN: "+format+"\n"+Reset, a...)
}

func LogInfo(format string, a ...any) {
	logImpl(Blue+"["+getNowTimeString()+"] INFO: "+format+"\n"+Reset, a...)
}

func LogSuccess(format string, a ...any) {
	logImpl(Green+"["+getNowTimeString()+"] INFO: "+format+"\n"+Reset, a...)
}

func getNowTimeString() string {
	return time.Now().String()
}

func logImpl(format string, a ...any) {
	_, err := fmt.Fprintf(os.Stdout, format, a...)
	if err != nil {
		return
	}
}
