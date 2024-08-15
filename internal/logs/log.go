package logs

import (
	"github.com/phuslu/log"
)

func Init() {
	log.DefaultLogger = log.Logger{
		Level:      log.InfoLevel,
		TimeFormat: "15:04:05",
		Caller:     1,
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		},
	}
}
