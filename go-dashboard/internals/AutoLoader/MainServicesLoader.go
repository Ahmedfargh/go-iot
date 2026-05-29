package autoloader

import (
	config "go-dashboard/internals/Config"
	models "go-dashboard/internals/Models"
	repository "go-dashboard/internals/Repository"
	auth_services "go-dashboard/internals/Services/Auth"
	crud_services "go-dashboard/internals/Services/Crud"
)

var Services_providers = make(map[string]interface{})

func LoadMainServices() {
	if Services_providers == nil {
		Services_providers = make(map[string]interface{})
	}

	db_object := config.DB
	admin_repo := repository.NewGormRepository[models.Admin](db_object)
	admin_service := crud_services.NewAdminService(admin_repo)
	Services_providers["admin.crud"] = admin_service

	admin_auth_service := auth_services.NewAdminAuthService(admin_repo)
	Services_providers["admin.auth"] = admin_auth_service

	role_repo := repository.NewGormRepository[models.Role](db_object)
	role_service := crud_services.NewRoleService(role_repo)
	Services_providers["role.crud"] = role_service

	permission_repo := repository.NewPermissionRepository(db_object)
	permission_service := auth_services.NewPermissionService(permission_repo)
	Services_providers["permission"] = permission_service

	device_repo := repository.NewGormRepository[models.Device](db_object)
	device_service := crud_services.NewDeviceService(device_repo)
	Services_providers["device.crud"] = device_service

	sector_repo := repository.NewGormRepository[models.Sector](db_object)
	sector_service := crud_services.NewSectorService(sector_repo)
	Services_providers["sector.crud"] = sector_service
}
