/*
Copyright 2023 XieYanke.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logging

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/hey-io/heycni/pkg/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once sync.Once
	log  zerolog.Logger
)

var DefaultLoggerOption = &LoggerOption{
	LogFile:    common.CNILogFile,
	LogLevel:   zerolog.InfoLevel,
	MaxSize:    100,
	MaxBackups: 10,
	MaxAge:     3,
	Compress:   true,
}

func InitLogger() *zerolog.Logger {
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

	return &log
}
