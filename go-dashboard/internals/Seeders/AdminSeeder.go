package seeders

import (
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
)

func SeedAdmins(adminService interfaces.ServiceInterface[models.Admin], permService interfaces.PermissionInterface) {
	admins := []struct {
		Name     string
		Email    string
		Phone    string
		Password string
		Role     string
	}{
		{
			Name:     "System Administrator",
			Email:    "admin@go-iot.com",
			Phone:    "0123456789",
			Password: "password123",
			Role:     "SuperAdmin",
		},
		{
			Name:     "John Engineer",
			Email:    "john@go-iot.com",
			Phone:    "0987654321",
			Password: "password123",
			Role:     "Engineer",
		},
	}

	for _, a := range admins {
		// Check if admin already exists
		existing, _, _ := adminService.GetAll(1, 100)
		found := false
		var currentAdmin *models.Admin
		for _, e := range existing {
			if e.Email == a.Email {
				found = true
				currentAdmin = &e
				break
			}
		}

		if !found {
			newAdmin := &models.Admin{
				Name:     a.Name,
				Email:    a.Email,
				Phone:    a.Phone,
				Password: a.Password,
			}
			var err error
			currentAdmin, err = adminService.Create(newAdmin)
			if err != nil {
				continue // Skip if creation failed
			}
		}

		// Assign role
		permService.AssignRole(currentAdmin, a.Role)
	}
}
