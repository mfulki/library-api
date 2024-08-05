package llog

import (
	"log"
	"os"
)

type defaultLogger struct {
	stdoutLog *log.Logger
}

func NewDefaultLogger() *defaultLogger {
	logger := &defaultLogger{
		stdoutLog: log.New(os.Stdout, "", 0),
	}

	return logger
}

func (l *defaultLogger) Print(vals ...any) {
	l.longLog(vals...)
}

func (l *defaultLogger) Fatal(vals ...any) {
	l.longLog(vals...)
	os.Exit(1)
}

func (l *defaultLogger) longLog(vals ...any) {
	title, _ := vals[0].(string)
	l.stdoutLog.Printf("\033[0;%dm%s\033[0m", color[title], title)
	l.stdoutLog.Printf("Time: %s", timeNow())

	if path := filePath(vals, 4); path != "" {
		l.stdoutLog.Printf("Path: %s", path)
	}

	l.stdoutLog.Println("Message: " + msg(vals[1:]))
}
