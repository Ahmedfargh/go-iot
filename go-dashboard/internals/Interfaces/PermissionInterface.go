package interfaces

import models "go-dashboard/internals/Models"

type PermissionInterface interface {
	// Role management
	AssignRole(admin *models.Admin, roleName string) error
	RemoveRole(admin *models.Admin, roleName string) error
	SyncRoles(admin *models.Admin, roleNames ...string) error
	HasRole(admin *models.Admin, roleName string) bool
	GetRoles(admin *models.Admin) ([]models.Role, error)

	// Permission management (Direct)
	GivePermissionTo(admin *models.Admin, permissionName string) error
	RevokePermissionTo(admin *models.Admin, permissionName string) error
	SyncPermissions(admin *models.Admin, permissionNames ...string) error
	HasDirectPermissionTo(admin *models.Admin, permissionName string) bool

	// Role-Permission management
	GivePermissionToRole(role *models.Role, permissionName string) error
	RevokePermissionFromRole(role *models.Role, permissionName string) error
	SyncPermissionsForRole(role *models.Role, permissionNames ...string) error

	// Comprehensive checks
	HasPermissionTo(admin *models.Admin, permissionName string) bool
}
