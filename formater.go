package slogger

import (
	"fmt"
	"time"
	"encoding/json"
	"strconv"
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
	return fmt.Sprintf("[%s] [%s]: %s", t, e.Level.String(), e.Message)
}

type JsonFormater struct {
}

func NewJsonFormater() *JsonFormater {
	return &JsonFormater{}
}

func (formater *JsonFormater) format(e *Event) string {
	type Out struct {
		TimeStamp int64  `json:"timeStamp"`
		Level     Level  `json:"level"`
		LevelDesc string `json:"LevelDesc"`
		Message   string `json:"message"`
	}

	o := Out{
		TimeStamp: time.Now().Unix(),
		Level:     e.Level,
		LevelDesc: e.Level.String(),
		Message:   e.Message,
	}

	j, _ := json.Marshal(o)
	return string(j)
}

type SeparationFormater struct {
	Delimiter string
}

func NewSeparationFormater(delimiter string) *SeparationFormater {
	return &SeparationFormater{
		Delimiter: delimiter,
	}
}

func (formater *SeparationFormater) format(e *Event) string {
	ret := ""
	ret += time.Now().Format("2006/1/2 15:04:05")
	ret += formater.Delimiter
	ret += strconv.Itoa(int(e.Level))
	ret += formater.Delimiter
	ret += e.Level.String()
	ret += formater.Delimiter
	ret += e.Message

	return ret
}
