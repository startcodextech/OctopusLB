package api

import (
	"github.com/phuslu/log"
	"github.com/startcodextech/managerlb/internal/config"
	"github.com/startcodextech/managerlb/pkg/system"
)

type Api struct {
	system  *system.SysInfo
	cfg     *config.Config
	modules map[string]Module
}

func Init() (*Api, error) {

	sys, err := system.Get()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get system info")
		return nil, err
	}

	cfg, err := config.Init()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create config")
		return nil, err
	}

	api := &Api{
		system:  sys,
		cfg:     cfg,
		modules: make(map[string]Module),
	}

	err = api.loadModules()
	if err != nil {
		log.Error().Err(err).Str("OctopusLB", "LoadModules").Msg("Failed")
		return nil, err
	}

	err = api.startModules()
	if err != nil {
		log.Error().Err(err).Str("OctopusLB", "StartModules").Msg("Failed")
		return nil, err
	}

	return api, nil
}
