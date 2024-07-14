package config

import "gorm.io/gorm"

type Certificates struct {
	gorm.Model
	Name        string
	Domains     string
	Email       string
	Environment string
	PathCert    string
	PathKey     string
}
