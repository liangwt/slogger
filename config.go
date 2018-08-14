package slogger

import "encoding/json"

func NewConfig(c []byte) (*Config, error) {
	config := new(Config)
	err := json.Unmarshal(c, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type LoggerConfig struct {
	Levels []string `json:"levels"`
}

type AppenderConfig struct {
	Out          string              `json:"out"`
	FileAppender *FileAppenderConfig `json:"fileAppender,omitempty"`
}

type FileAppenderConfig struct {
	FileName *struct {
		TRACE string `json:",omitempty"`
		DEBUG string `json:",omitempty"`
		INFO  string `json:",omitempty"`
		WARN  string `json:",omitempty"`
		ERROR string `json:",omitempty"`
		FATAL string `json:",omitempty"`
		ALL   string `json:",omitempty"`
	} `json:"fileName,omitempty"`
}

type FormaterConfig struct {
	Format string `json:"format,omitempty"`
}

type Config struct {
	Logger   *LoggerConfig   `json:"logger"`
	Appender *AppenderConfig `json:"appender"`
	Formater *FormaterConfig `json:"formater"`
}
