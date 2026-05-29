package seeders

import (
	models "go-dashboard/internals/Models"

	"gorm.io/gorm"
)

func SeedPermissions(db *gorm.DB) {
	permissions := []models.Permission{
		{Name: "view-dashboard", Description: "Can view the main dashboard overview"},
		{Name: "manage-admins", Description: "Can create, update and delete admin accounts"},
		{Name: "manage-roles", Description: "Can manage roles and their permissions"},
		{Name: "manage-iot-devices", Description: "Can manage IoT device configurations"},
		{Name: "view-system-logs", Description: "Can view system logs and activities"},
	}

	for _, p := range permissions {
		// Use FirstOrCreate to avoid duplicates
		db.Where(models.Permission{Name: p.Name}).FirstOrCreate(&p)
	}
}
