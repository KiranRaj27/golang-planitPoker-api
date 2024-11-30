package logger

import (
	"log"
	"os"
)

type Logger struct {
    infoLog  *log.Logger
    errorLog *log.Logger
}

func NewLogger() *Logger {
    return &Logger{
        infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
        errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}

func (l *Logger) Info(message string) {
    l.infoLog.Println(message)
}

func (l *Logger) Error(err error) {
    l.errorLog.Println(err)
}
