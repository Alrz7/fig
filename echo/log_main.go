package echo

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

var (
	DefultInfoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	DefultErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)

type Logger struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
}

var DefultLogger = Logger{
	ErrorLogger: DefultErrorLog,
	InfoLogger:  DefultInfoLog,
}

func (l *Logger) Errort(text string, a ...any) {
	txt := fmt.Sprintf(text, a...)
	trace := fmt.Sprintf("%v: %s", txt, debug.Stack())
	l.ErrorLogger.Output(2, trace)
	os.Exit(1)
}

func (l *Logger) Error(err error, text string, a ...any) {
	if err != nil {
		txt := fmt.Sprintf(text, a...)
		trace := fmt.Sprintf("%v: %s\n%s", txt, error.Error(err), debug.Stack())
		l.ErrorLogger.Output(2, trace)
		os.Exit(1)
	}
}

func (l *Logger) Info(text string) {
	l.InfoLogger.Println(text)
}

// 	stac := trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
