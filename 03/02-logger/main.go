package main

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	*log.Logger
	level LogLevel
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.level >= srcLogLvl {
		l.Logger.Println(prefix + " " + msg)
	}
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	l.level = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO", msg)
}

func (l *LogExtended) Warnln(msg string) {

	l.println(LogLevelWarning, "WARN", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERROR", msg)
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
