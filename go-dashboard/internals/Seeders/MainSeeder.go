package seeders

import (
	"log"

	config "go-dashboard/internals/Config"
	models "go-dashboard/internals/Models"
	repository "go-dashboard/internals/Repository"
	auth_services "go-dashboard/internals/Services/Auth"
	crud_services "go-dashboard/internals/Services/Crud"
)

func SeedAll() {
	db := config.DB

	// Repositories
	adminRepo := repository.NewGormRepository[models.Admin](db)
	roleRepo := repository.NewGormRepository[models.Role](db)
	permRepo := repository.NewPermissionRepository(db)
	sectorRepo := repository.NewGormRepository[models.Sector](db)

	// Services
	adminService := crud_services.NewAdminService(adminRepo)
	roleService := crud_services.NewRoleService(roleRepo)
	permissionService := auth_services.NewPermissionService(permRepo)
	sectorService := crud_services.NewSectorService(sectorRepo)

	log.Println("Seeding Permissions...")
	SeedPermissions(db)

	log.Println("Seeding Roles...")
	SeedRoles(db, roleService, permissionService)

	log.Println("Seeding Admins...")
	SeedAdmins(adminService, permissionService)

	log.Println("Seeding Sectors...")
	SeedSectors(sectorService)

	log.Println("Seeding completed successfully.")
}
