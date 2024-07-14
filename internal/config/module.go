package config

import (
	"errors"
	"github.com/phuslu/log"
	"gorm.io/gorm"
)

type Module struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex;not null"`
	DefaultSystem bool   `gorm:"default:false;not null"`
	Installed     bool   `gorm:"default:false;not null"`
	Started       bool   `gorm:"default:false;not null"`
	Enabled       bool   `gorm:"default:false;not null"`
	Tool          string
	ToolVersion   string
}

func (c *Config) LoadModules() ([]Module, error) {
	var modules []Module
	result := c.db.Find(&modules)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to load modules")
		return nil, result.Error
	}
	return modules, nil
}

func (c *Config) AddModule(module *Module) error {
	result := c.db.Create(&module)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to add module")
		return result.Error
	}
	return nil
}

func (c *Config) UpdateModule(module *Module) error {
	result := c.db.Save(&module)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update module")
		return result.Error
	}
	return nil
}

func (c *Config) DeleteModule(module *Module) error {
	result := c.db.Delete(&module)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to delete module")
		return result.Error
	}
	return nil
}

func (c *Config) GetModuleByName(name string) (*Module, error) {
	var module Module
	result := c.db.First(&module, "name = ?", name)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to get module")
		return nil, result.Error
	}
	return &module, nil
}

func (c *Config) RegistreModule(defaultValues *Module) (*Module, error) {
	existingModule, err := c.GetModuleByName(defaultValues.Name)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("Failed register module")
			return nil, err
		}
		if existingModule == nil {
			err = c.AddModule(defaultValues)
			if err != nil {
				log.Error().Err(err).Msg("Failed register module")
				return nil, err
			}
			log.Info().Msg("Successfully registered module: " + defaultValues.Name)
			return defaultValues, nil
		}
	}

	return existingModule, nil
}
