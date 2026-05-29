package seeders

import (
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB, roleService interfaces.ServiceInterface[models.Role], permService interfaces.PermissionInterface) {
	roles := []struct {
		Name        string
		Description string
		Permissions []string
	}{
		{
			Name:        "SuperAdmin",
			Description: "Full access to all system features",
			Permissions: []string{"view-dashboard", "manage-admins", "manage-roles", "manage-iot-devices", "view-system-logs"},
		},
		{
			Name:        "Engineer",
			Description: "Technical management of IoT devices",
			Permissions: []string{"view-dashboard", "manage-iot-devices"},
		},
		{
			Name:        "Viewer",
			Description: "ReadOnly access to the dashboard",
			Permissions: []string{"view-dashboard"},
		},
	}

	for _, r := range roles {
		var role models.Role
		// Use direct DB for existence check to avoid service layer constraints if any
		err := db.Where("name = ?", r.Name).First(&role).Error
		if err != nil {
			// Role doesn't exist, create it
			role = models.Role{
				Name:        r.Name,
				Description: r.Description,
			}
			db.Create(&role)
		}

		// Sync permissions
		permService.SyncPermissionsForRole(&role, r.Permissions...)
	}
}
