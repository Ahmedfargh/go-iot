package config

import (
	models "go-dashboard/internals/Models"
)

func MigrateModels() {
	connection := DB
	connection.AutoMigrate(&models.Admin{}, &models.Role{}, &models.Permission{}, &models.Device{}, &models.Sector{})
}
