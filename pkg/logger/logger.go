package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once sync.Once
	log  zerolog.Logger
)

var DefaultLoggerOption = &LoggerOption{
	LogFile:    "hey-cni.log", //common.DefaultLogPath,
	LogLevel:   zerolog.InfoLevel,
	MaxSize:    100,
	MaxBackups: 10,
	MaxAge:     3,
	Compress:   true,
}

func InitLogger() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano
		var output io.Writer

		_, exist := os.LookupEnv("HEY_CNI_DEBUG")
		if exist {
			output = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			}
		} else {
			fileLogger := &lumberjack.Logger{
				Filename:   DefaultLoggerOption.LogFile,
				MaxSize:    DefaultLoggerOption.MaxSize, // megabytes
				MaxBackups: DefaultLoggerOption.MaxBackups,
				MaxAge:     DefaultLoggerOption.MaxAge, // days
				Compress:   DefaultLoggerOption.Compress,
			}

			output = zerolog.MultiLevelWriter(os.Stdout, fileLogger)
		}

		log = zerolog.New(output).
			Level(DefaultLoggerOption.LogLevel).
			With().
			Timestamp().
			Str("scope", "default").
			Logger()
	})

	return log
}
