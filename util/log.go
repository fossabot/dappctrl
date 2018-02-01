package util

import (
	"log"
	"os"
)

type Logger log.Logger

func NewLogger() *Logger {
	return (*Logger)(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile))
}

const (
	LOG_INFO  = iota
	LOG_WARN  = iota
	LOG_ERROR = iota
	LOG_FATAL = iota
)

func (l *Logger) Log(lvl int, fmt string, v ...interface{}) {
	var prefix string

	switch lvl {
	case LOG_INFO:
		prefix = "INFO "
	case LOG_WARN:
		prefix = "WARNING "
	case LOG_ERROR:
		prefix = "ERROR "
	case LOG_FATAL:
		prefix = "FATAL "
	}

	(*log.Logger)(l).Printf(prefix+fmt, v...)

	if lvl == LOG_FATAL {
		os.Exit(1)
	}
}

func (l *Logger) Info(fmt string, v ...interface{}) {
	l.Log(LOG_INFO, fmt, v...)
}

func (l *Logger) Warn(fmt string, v ...interface{}) {
	l.Log(LOG_WARN, fmt, v...)
}

func (l *Logger) Error(fmt string, v ...interface{}) {
	l.Log(LOG_ERROR, fmt, v...)
}

func (l *Logger) Fatal(fmt string, v ...interface{}) {
	l.Log(LOG_FATAL, fmt, v...)
}
