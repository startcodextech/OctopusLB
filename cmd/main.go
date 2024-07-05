package main

import (
	"github.com/phuslu/log"
	"github.com/startcodextech/managerlb/system"
)

func init() {
	log.DefaultLogger = log.Logger{
		Level:      log.DebugLevel,
		TimeFormat: "15:04:05",
		Caller:     1,
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		},
	}
}

func main() {
	log.Info().Msg("Starting OpenLB")

	sys := system.Init()

	err := sys.OS.InstallPackage("nginx")
	if err != nil {
		log.Error().Err(err).Msg("Failed to install package")
	}

	err = sys.OS.UninstallPackage("nginx")
	if err != nil {
		log.Error().Err(err).Msg("Failed to uninstall package")
	}

	select {}
}
