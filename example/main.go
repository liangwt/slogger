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
