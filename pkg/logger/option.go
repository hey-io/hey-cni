package logger

import "github.com/rs/zerolog"

type LoggerOption struct {
	LogFile    string
	LogLevel   zerolog.Level
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}
