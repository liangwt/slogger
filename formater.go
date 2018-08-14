package slogger

import (
	"fmt"
	"time"
)

type Formater interface {
	format(e *Event) string
}

type DefaultFormater struct {
}

func NewDefaultFormater() *DefaultFormater {
	return &DefaultFormater{}
}

func (formater *DefaultFormater) format(e *Event) string {
	t := time.Now().Format("2006/1/2 15:04:05")
	return fmt.Sprintf("[%s] [%s]: %s \n", t, e.Level.String(), e.Message)
}
