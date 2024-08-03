package llog

import (
	"fmt"
	"log"
	"os"
)

type FileLogger struct {
	stdoutLog *log.Logger
	fileLog   *os.File
}

func NewFileLogger(fileName string) (*FileLogger, error) {
	filePath := fmt.Sprintf("logs/%s.log", fileName)
	fileLog, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := &FileLogger{
		stdoutLog: log.New(fileLog, "", 0),
		fileLog:   fileLog,
	}

	return logger, nil
}

func (l *FileLogger) Print(vals ...any) {
	l.stdoutLog.Print(vals...)
}

func (l *FileLogger) Fatal(vals ...any) {
	l.stdoutLog.Fatal(vals...)
	os.Exit(1)
}

func (l *FileLogger) Close(vals ...any) {
	if l.fileLog != nil {
		l.fileLog.Close()
	}
}
