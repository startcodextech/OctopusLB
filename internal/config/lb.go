package config

import (
	"github.com/phuslu/log"
	"gorm.io/gorm"
)

type (
	LbApp struct {
		gorm.Model
		App        string `json:"app"`
		IsSSL      bool   `json:"is_ssl"`
		Host       string `json:"host" gorm:"uniqueIndex"`
		IsTCP      bool   `json:"is_tcp"`
		Port       int    `json:"port"`
		RemotePort int    `json:"remote_port"`
		RemoteHost string `json:"remote_host"`
		Enable     bool   `json:"enable"`
		ModuleID   uint
		Module     Module
	}
)

func (c *Config) AddLbApp(lbApp *LbApp) error {
	result := c.db.Create(&lbApp)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to add lb app")
		return result.Error
	}
	return nil
}

func (c *Config) UpdateLbApp(lbApp *LbApp) error {
	result := c.db.Save(&lbApp)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to update lb app")
		return result.Error
	}
	return nil
}

func (c *Config) DeleteLbApp(lbApp *LbApp) error {
	result := c.db.Delete(&lbApp)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to delete lb app")
		return result.Error
	}
	return nil
}
