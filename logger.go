package slogger

import "fmt"

type Logger struct {
	appender *Appender
	levels   []Level
}

func NewLogger(appender *Appender, levels []Level) *Logger {
	logger := new(Logger)
	logger.appender = appender
	logger.levels = levels
	return logger
}

func (logger *Logger) SetAppender(appender *Appender) {
	logger.appender = appender
}

func (logger *Logger) TRACE(format string, data ...interface{}) {
	if inLevels(TRACE, logger.levels) {
		msg := fmt.Sprintf(format, data...)
		e := NewEvent(TRACE, msg)
		(*logger.appender).out(e)
	}
}

func (logger *Logger) DEBUG(format string, data ...interface{}) {
	if inLevels(DEBUG, logger.levels) {
		msg := fmt.Sprintf(format, data...)
		e := NewEvent(DEBUG, msg)
		(*logger.appender).out(e)
	}
}

func (logger *Logger) INFO(format string, data ...interface{}) {
	if inLevels(INFO, logger.levels) {
		msg := fmt.Sprintf(format, data...)
		e := NewEvent(INFO, msg)
		(*logger.appender).out(e)
	}
}

func (logger *Logger) WARN(format string, data ...interface{}) {
	if inLevels(WARN, logger.levels) {
		msg := fmt.Sprintf(format, data...)
		e := NewEvent(WARN, msg)
		(*logger.appender).out(e)
	}
}

func (logger *Logger) ERROR(format string, data ...interface{}) {
	if inLevels(ERROR, logger.levels) {
		msg := fmt.Sprintf(format, data...)
		e := NewEvent(ERROR, msg)
		(*logger.appender).out(e)
	}
}

func (logger *Logger) FATAL(format string, data ...interface{}) {
	if inLevels(FATAL, logger.levels) {
		msg := fmt.Sprintf(format, data...)
		e := NewEvent(FATAL, msg)
		(*logger.appender).out(e)
	}
}

func inLevels(l Level, levels []Level) bool {
	for _, v := range levels {
		if v == l {
			return true
		}
	}

	return false
}
