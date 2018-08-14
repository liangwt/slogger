package slogger

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	ALL
)

type Level int

func NewLevel(s string) Level {
	var l Level

	switch s {
	case "TRACE":
		l = TRACE
	case "DEBUG":
		l = DEBUG
	case "INFO":
		l = INFO
	case "WARN":
		l = WARN
	case "ERROR":
		l = ERROR
	case "FATAL":
		l = FATAL
	case "ALL":
		l = ALL
	}

	return l
}

func (l Level) String() string {
	var s string

	switch l {
	case TRACE:
		s = "TRACE"
	case DEBUG:
		s = "DEBUG"
	case INFO:
		s = "INFO"
	case WARN:
		s = "WARN"
	case ERROR:
		s = "ERROR"
	case FATAL:
		s = "FATAL"
	case ALL:
		s = "ALL"
	}

	return s
}

type Event struct {
	Level   Level
	Message string
}

func NewEvent(level Level, message string) *Event {
	return &Event{
		Level:   level,
		Message: message,
	}
}



