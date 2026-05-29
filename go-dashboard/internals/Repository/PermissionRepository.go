package repository

import (
	models "go-dashboard/internals/Models"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) FindRoleByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

func (r *PermissionRepository) FindPermissionByName(name string) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.Where("name = ?", name).First(&permission).Error
	return &permission, err
}

func (r *PermissionRepository) AssignRole(admin *models.Admin, role *models.Role) error {
	return r.db.Model(admin).Association("Roles").Append(role)
}

func (r *PermissionRepository) RemoveRole(admin *models.Admin, role *models.Role) error {
	return r.db.Model(admin).Association("Roles").Delete(role)
}

func (r *PermissionRepository) SyncRoles(admin *models.Admin, roles []models.Role) error {
	return r.db.Model(admin).Association("Roles").Replace(roles)
}

func (r *PermissionRepository) GetRoles(admin *models.Admin) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Model(admin).Association("Roles").Find(&roles)
	return roles, err
}

func (r *PermissionRepository) GivePermissionTo(admin *models.Admin, permission *models.Permission) error {
	return r.db.Model(admin).Association("Permissions").Append(permission)
}

func (r *PermissionRepository) RevokePermissionTo(admin *models.Admin, permission *models.Permission) error {
	return r.db.Model(admin).Association("Permissions").Delete(permission)
}

func (r *PermissionRepository) SyncPermissions(admin *models.Admin, permissions []models.Permission) error {
	return r.db.Model(admin).Association("Permissions").Replace(permissions)
}

func (r *PermissionRepository) GivePermissionToRole(role *models.Role, permission *models.Permission) error {
	return r.db.Model(role).Association("Permissions").Append(permission)
}

func (r *PermissionRepository) RevokePermissionFromRole(role *models.Role, permission *models.Permission) error {
	return r.db.Model(role).Association("Permissions").Delete(permission)
}

func (r *PermissionRepository) SyncPermissionsForRole(role *models.Role, permissions []models.Permission) error {
	return r.db.Model(role).Association("Permissions").Replace(permissions)
}

func (r *PermissionRepository) GetAdminPermissions(admin *models.Admin) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Model(admin).Association("Permissions").Find(&permissions)
	return permissions, err
}

func (r *PermissionRepository) GetRolePermissions(role *models.Role) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Model(role).Association("Permissions").Find(&permissions)
	return permissions, err
}
