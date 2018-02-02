package util

import (
	"errors"
	"log"
	"os"
	"strings"
)

type Logger struct {
	logger *log.Logger
	level  int
}

type LogConfig struct {
	Level string
}

const (
	LOG_DEBUG   = iota
	LOG_INFO    = iota
	LOG_WARNING = iota
	LOG_ERROR   = iota
	LOG_FATAL   = iota
)

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level: logLevelStrs[LOG_INFO],
	}
}

var logLevelStrs = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}

func parseLogLevel(lvl string) int {
	switch strings.ToUpper(lvl) {
	case logLevelStrs[LOG_DEBUG]:
		return LOG_DEBUG
	case logLevelStrs[LOG_INFO]:
		return LOG_INFO
	case logLevelStrs[LOG_WARNING]:
		return LOG_WARNING
	case logLevelStrs[LOG_ERROR]:
		return LOG_ERROR
	case logLevelStrs[LOG_FATAL]:
		return LOG_FATAL
	}

	return -1
}

func NewLogger(conf *LogConfig) (*Logger, error) {
	lvl := parseLogLevel(conf.Level)
	if lvl < LOG_DEBUG || lvl > LOG_FATAL {
		return nil, errors.New("bad log level")
	}

	return &Logger{
		log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile),
		lvl,
	}, nil
}

func (l *Logger) Log(lvl int, fmt string, v ...interface{}) {
	if lvl < LOG_DEBUG || lvl > LOG_FATAL || lvl < l.level {
		return
	}

	l.logger.Printf(logLevelStrs[lvl]+" "+fmt, v...)

	if lvl == LOG_FATAL {
		os.Exit(1)
	}
}

func (l *Logger) Debug(fmt string, v ...interface{}) {
	l.Log(LOG_DEBUG, fmt, v...)
}

func (l *Logger) Info(fmt string, v ...interface{}) {
	l.Log(LOG_INFO, fmt, v...)
}

func (l *Logger) Warn(fmt string, v ...interface{}) {
	l.Log(LOG_WARNING, fmt, v...)
}

func (l *Logger) Error(fmt string, v ...interface{}) {
	l.Log(LOG_ERROR, fmt, v...)
}

func (l *Logger) Fatal(fmt string, v ...interface{}) {
	l.Log(LOG_FATAL, fmt, v...)
}
