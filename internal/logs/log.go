package logs

import (
	"github.com/phuslu/log"
	"os"
	"strings"
)

func Init() {
	var level log.Level

	logLevel := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	switch logLevel {
	case "DEBUG":
		level = log.DebugLevel
	case "INFO":
		level = log.InfoLevel
	case "WARN":
		level = log.WarnLevel
	case "ERROR":
		level = log.ErrorLevel
	case "FATAL":
		level = log.FatalLevel
	case "PANIC":
		level = log.PanicLevel
	default:
		level = log.TraceLevel
	}

	log.DefaultLogger = log.Logger{
		Level:      level,
		TimeFormat: "15:04:05",
		Caller:     1,
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		},
	}
}
