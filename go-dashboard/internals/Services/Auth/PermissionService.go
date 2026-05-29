package auth

import (
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
	repository "go-dashboard/internals/Repository"
)

type PermissionService struct {
	repo *repository.PermissionRepository
}

func NewPermissionService(repo *repository.PermissionRepository) interfaces.PermissionInterface {
	return &PermissionService{repo: repo}
}

func (s *PermissionService) AssignRole(admin *models.Admin, roleName string) error {
	role, err := s.repo.FindRoleByName(roleName)
	if err != nil {
		return err
	}
	return s.repo.AssignRole(admin, role)
}

func (s *PermissionService) RemoveRole(admin *models.Admin, roleName string) error {
	role, err := s.repo.FindRoleByName(roleName)
	if err != nil {
		return err
	}
	return s.repo.RemoveRole(admin, role)
}

func (s *PermissionService) SyncRoles(admin *models.Admin, roleNames ...string) error {
	var roles []models.Role
	for _, name := range roleNames {
		role, err := s.repo.FindRoleByName(name)
		if err != nil {
			return err
		}
		roles = append(roles, *role)
	}
	return s.repo.SyncRoles(admin, roles)
}

func (s *PermissionService) HasRole(admin *models.Admin, roleName string) bool {
	roles, err := s.repo.GetRoles(admin)
	if err != nil {
		return false
	}
	for _, r := range roles {
		if r.Name == roleName {
			return true
		}
	}
	return false
}

func (s *PermissionService) GetRoles(admin *models.Admin) ([]models.Role, error) {
	return s.repo.GetRoles(admin)
}

func (s *PermissionService) GivePermissionTo(admin *models.Admin, permissionName string) error {
	permission, err := s.repo.FindPermissionByName(permissionName)
	if err != nil {
		return err
	}
	return s.repo.GivePermissionTo(admin, permission)
}

func (s *PermissionService) RevokePermissionTo(admin *models.Admin, permissionName string) error {
	permission, err := s.repo.FindPermissionByName(permissionName)
	if err != nil {
		return err
	}
	return s.repo.RevokePermissionTo(admin, permission)
}

func (s *PermissionService) SyncPermissions(admin *models.Admin, permissionNames ...string) error {
	var permissions []models.Permission
	for _, name := range permissionNames {
		permission, err := s.repo.FindPermissionByName(name)
		if err != nil {
			return err
		}
		permissions = append(permissions, *permission)
	}
	return s.repo.SyncPermissions(admin, permissions)
}

func (s *PermissionService) HasDirectPermissionTo(admin *models.Admin, permissionName string) bool {
	permissions, err := s.repo.GetAdminPermissions(admin)
	if err != nil {
		return false
	}
	for _, p := range permissions {
		if p.Name == permissionName {
			return true
		}
	}
	return false
}

func (s *PermissionService) GivePermissionToRole(role *models.Role, permissionName string) error {
	permission, err := s.repo.FindPermissionByName(permissionName)
	if err != nil {
		return err
	}
	return s.repo.GivePermissionToRole(role, permission)
}

func (s *PermissionService) RevokePermissionFromRole(role *models.Role, permissionName string) error {
	permission, err := s.repo.FindPermissionByName(permissionName)
	if err != nil {
		return err
	}
	return s.repo.RevokePermissionFromRole(role, permission)
}

func (s *PermissionService) SyncPermissionsForRole(role *models.Role, permissionNames ...string) error {
	var permissions []models.Permission
	for _, name := range permissionNames {
		permission, err := s.repo.FindPermissionByName(name)
		if err != nil {
			return err
		}
		permissions = append(permissions, *permission)
	}
	return s.repo.SyncPermissionsForRole(role, permissions)
}

func (s *PermissionService) HasPermissionTo(admin *models.Admin, permissionName string) bool {
	// Check direct permissions first
	if s.HasDirectPermissionTo(admin, permissionName) {
		return true
	}

	// Check permissions via roles
	roles, err := s.repo.GetRoles(admin)
	if err != nil {
		return false
	}

	for _, role := range roles {
		permissions, err := s.repo.GetRolePermissions(&role)
		if err != nil {
			continue
		}
		for _, p := range permissions {
			if p.Name == permissionName {
				return true
			}
		}
	}

	return false
}
