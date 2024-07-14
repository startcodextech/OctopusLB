package api

import (
	"errors"
	"github.com/phuslu/log"
	"github.com/startcodextech/octopuslb/internal/lb"
)

var (
	ErrorModuleNotFound = errors.New("module not found")
)

type Module interface {
	Name() string
	Install() error
	Uninstall() error
	Start() error
	Stop() error
	Reload() error
	IsInstalled() bool
	IsEnabled() bool
	IsDefaultSystem() bool
}

func (api *Api) loadModules() error {
	log.Info().Str("OctopusLB", "LoadModules").Msg("...")

	// Module Load Balancer
	loadBalancer, err := lb.New(api.system.PlatformFamily, api.cfg)
	if err != nil {
		log.Error().Str("OctopusLB", "LoadBalancer module").Msg("Failed")
		return err
	}
	api.modules[loadBalancer.Name()] = loadBalancer
	log.Info().Str("OctopusLB", "LoadBalancer").Msg("OK")

	log.Info().Str("OctopusLB", "LoadModules").Msg("OK")

	return nil
}

func (api *Api) startModules() error {
	log.Info().Str("OctopusLB", "StartModules").Msg("...")

	for _, module := range api.modules {
		if module.IsDefaultSystem() || module.IsEnabled() {
			if !module.IsInstalled() {
				err := module.Install()
				if err != nil {
					log.Error().
						Err(err).
						Str("OctopusLB", "InstallModules").
						Str("Name", module.Name()).
						Msg("Failed")
					return err
				}
			}
			err := module.Start()
			if err != nil {
				log.Error().
					Err(err).
					Str("OctopusLB", "StartModules").
					Str("Name", module.Name()).
					Msg("Failed")
				return err
			}
			log.Info().
				Str("OctopusLB", "StartModules").
				Str("Name", module.Name()).
				Msg("OK")
		}
	}

	log.Info().Str("OctopusLB", "StartModules").Msg("OK")
	return nil
}

func (api *Api) InstallModule(name string) error {
	module, err := api.getModule(name)
	if err != nil {
		return err
	}
	err = module.Install()

	if err != nil {
		log.Error().
			Err(err).
			Str("OctopusLB", "InstallModule").
			Str("Name", name).
			Msg("Failed")
		return err
	}

	log.Info().
		Str("OctopusLB", "InstallModule").
		Str("Name", name).
		Msg("OK")
	return nil
}

func (api *Api) UninstallModule(name string) error {
	module, err := api.getModule(name)
	if err != nil {
		return err
	}
	err = module.Uninstall()

	if err != nil {
		log.Error().
			Err(err).
			Str("OctopusLB", "UninstallModule").
			Str("Name", name).
			Msg("Failed")
		return err
	}

	log.Info().
		Str("OctopusLB", "UninstallModule").
		Str("Name", name).
		Msg("OK")
	return nil
}

func (api *Api) StartModule(name string) error {
	module, err := api.getModule(name)
	if err != nil {
		return err
	}
	err = module.Start()

	if err != nil {
		log.Error().
			Err(err).
			Str("OctopusLB", "StartModule").
			Str("Name", name).
			Msg("Failed")
		return err
	}

	log.Info().
		Str("OctopusLB", "StartModule").
		Str("Name", name).
		Msg("OK")
	return nil
}

func (api *Api) StopModule(name string) error {
	module, err := api.getModule(name)
	if err != nil {
		return err
	}
	err = module.Stop()

	if err != nil {
		log.Error().
			Err(err).
			Str("OctopusLB", "StopModule").
			Str("Name", name).
			Msg("Failed")
		return err
	}

	log.Info().
		Str("OctopusLB", "StopModule").
		Str("Name", name).
		Msg("OK")
	return nil
}

func (api *Api) getModule(name string) (Module, error) {
	module, ok := api.modules[name]
	if !ok {
		log.Error().
			Err(ErrorModuleNotFound).
			Str("OctopusLB", "FindModule").
			Str("Name", name).
			Msg("Failed")
		return nil, ErrorModuleNotFound
	}
	return module, nil
}
