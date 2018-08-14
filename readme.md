## A Simaple Event Log Tool - Slogger

### Introduction

Easy-to-use event log system that can implement output specific event to a file or console with a simple configuration while meeting custom formats

### Features

- Simaple configuration (default parse json)
- Six event support(TRACE DEBUG INFO WARN ERROR FATAL)
- Specific event level output define by configure
- File or console output implement
- Different events output to different files
- Custom format

### Usages

#### simple usage

```go
package main

import (
	"github.com/liangwt/slogger"
)

func main() {
	c := `
	{
  		"logger": {
    		"levels": ["DEBUG", "ERROR", "INFO"]
  		},
		"appender": {
			"out": "file",
			"fileAppender":{
				"filename": {
					"ALL": "./example/ALL_%T[20060102_15].log",
					"ERROR": "./example/ERROR_%T[20060102_15]_custom.log"
				}
			}
		},
		"formater": {
			"format": "default"
		}
	}
	`
	config, _ := slogger.NewConfig([]byte(c))

	logger := slogger.InitLogger(config)
	logger.ERROR("this is an error message: %s", "ERROR")
	logger.INFO("this is an info message: %s", "INFO")
	logger.TRACE("this is an trace message: %s", "TRACE")
	logger.DEBUG("this is an debug message: %s", "DEBUG")
}
```
#### fileAppender configure

- %E: Write different files for different events
- %T[20060102_15]: Use the time formatting method in golang
- "appender.fileAppender.filename.ALL" field
  
  If only the `appender.fileAppender.filename.ALL` field is configured and this field does not contain "%E" in the value, all events specified in `logger.levels` will be logged to the same file.

  Otherside, if the `appender.fileAppender.filename.ALL` field contains "%E", all events specified in `logger.levels` will be logged to the different file named by event and also will be logged to a file named "ALL".

  Note that if you not only set "ALL" field but also set each event, the latter will overwrite the "ALL".

### Todo
- http appender
- log analyse and show
- more formatter
- test file
