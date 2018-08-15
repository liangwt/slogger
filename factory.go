package slogger

import (
	"regexp"
	"time"
	"strings"
)

func InitLogger(conf *Config) *Logger {
	// init formater
	var formater Formater
	switch conf.Formater.Format {
	case "json":
		formater = NewJsonFormater()
	case "separation":
		formater = NewSeparationFormater(conf.Formater.SeparationFormater.Delimiter)
	case "default":
		formater = NewDefaultFormater()
	default:
		formater = NewDefaultFormater()
	}

	// init appender
	var appender Appender
	switch conf.Appender.Out {
	case "file":
		fileNames := make(map[Level]string, 7)
		if conf.Appender.FileAppender.FileName.ALL != "" {
			if strings.Contains(conf.Appender.FileAppender.FileName.ALL, "%E") {
				fileNames[ALL] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, ALL)
				fileNames[TRACE] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, TRACE)
				fileNames[DEBUG] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, DEBUG)
				fileNames[INFO] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, INFO)
				fileNames[WARN] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, WARN)
				fileNames[ERROR] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, ERROR)
				fileNames[FATAL] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, FATAL)
			} else {
				fileNames[ALL] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, ALL)
			}
		}
		if conf.Appender.FileAppender.FileName.TRACE != "" {
			fileNames[TRACE] = FormatFileName(conf.Appender.FileAppender.FileName.ALL, TRACE)
		}
		if conf.Appender.FileAppender.FileName.DEBUG != "" {
			fileNames[DEBUG] = FormatFileName(conf.Appender.FileAppender.FileName.DEBUG, DEBUG)
		}
		if conf.Appender.FileAppender.FileName.INFO != "" {
			fileNames[INFO] = FormatFileName(conf.Appender.FileAppender.FileName.INFO, INFO)
		}
		if conf.Appender.FileAppender.FileName.WARN != "" {
			fileNames[WARN] = FormatFileName(conf.Appender.FileAppender.FileName.WARN, WARN)
		}
		if conf.Appender.FileAppender.FileName.ERROR != "" {
			fileNames[ERROR] = FormatFileName(conf.Appender.FileAppender.FileName.ERROR, ERROR)
		}
		if conf.Appender.FileAppender.FileName.FATAL != "" {
			fileNames[FATAL] = FormatFileName(conf.Appender.FileAppender.FileName.FATAL, FATAL)
		}
		appender = NewFileAppender(&formater, fileNames)
	case "console":
		appender = NewConsoleAppender(&formater)
	default:
		appender = NewConsoleAppender(&formater)
	}

	// init logger
	var logger *Logger
	levels := make([]Level, len(conf.Logger.Levels), len(conf.Logger.Levels))
	for i, level := range conf.Logger.Levels {
		levels[i] = NewLevel(level)
	}
	logger = NewLogger(&appender, levels)

	return logger
}

func FormatFileName(fileNameFmt string, l Level) string {
	timeExpr := `%T\[(?P<layout>.*)\]`
	timeRe := regexp.MustCompile(timeExpr)

	var afterTime string
	layout := timeRe.FindAllStringSubmatch(fileNameFmt, -1)
	if len(layout) == 0 {
		afterTime = fileNameFmt
	} else {
		t := time.Now().Format(layout[0][1])
		afterTime = timeRe.ReplaceAllString(fileNameFmt, t)
	}

	eventExpr := `%E`
	eventRe := regexp.MustCompile(eventExpr)
	filename := eventRe.ReplaceAllString(afterTime, l.String())

	return filename
}
