package slogger

import (
	"os"
)

type Appender interface {
	out(e *Event)
	SetFormater(formater *Formater)
}

type ConsoleAppender struct {
	formater *Formater
}

func NewConsoleAppender(formater *Formater) *ConsoleAppender {
	return &ConsoleAppender{
		formater: formater,
	}
}

func (appender *ConsoleAppender) out(e *Event) {
	formatString := (*appender.formater).format(e)
	os.Stdout.WriteString(formatString)
}

func (appender *ConsoleAppender) SetFormater(formater *Formater) {
	appender.formater = formater
}

type FileAppender struct {
	formater  *Formater
	fileNames map[Level]string
}

func NewFileAppender(formater *Formater, fileNames map[Level]string) *FileAppender {
	return &FileAppender{
		formater:  formater,
		fileNames: fileNames,
	}
}

func (appender *FileAppender) out(e *Event) {
	if v, ok := appender.fileNames[ALL]; ok {
		f, _ := os.OpenFile(v, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		f.WriteString((*appender.formater).format(e))
		defer f.Close()
	}
	if v, ok := appender.fileNames[e.Level]; ok {
		f, _ := os.OpenFile(v, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		f.WriteString((*appender.formater).format(e))
		defer f.Close()
	}
}

func (appender *FileAppender) SetFormater(formater *Formater) {
	appender.formater = formater
}
