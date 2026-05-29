package crud

import (
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
)

type RoleService struct {
	repo interfaces.RepositoryInterface[models.Role]
}

func NewRoleService(repo interfaces.RepositoryInterface[models.Role]) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) Create(role *models.Role) (*models.Role, error) {
	return s.repo.Create(role)
}

func (s *RoleService) GetByID(id uint) (*models.Role, error) {
	return s.repo.With("Permissions").FindById(id)
}

func (s *RoleService) GetAll(page, pageSize int) ([]models.Role, int64, error) {
	count, _ := s.repo.Count(map[string]interface{}{})
	results, err := s.repo.With("Permissions").Paginate(page, pageSize).FindByWhere(map[string]interface{}{})
	return results, count, err
}

func (s *RoleService) Update(id uint, updatedData *models.Role) (*models.Role, error) {
	current, err := s.repo.FindById(id)
	if err != nil || current == nil {
		return nil, err
	}
	current.Name = updatedData.Name
	current.Description = updatedData.Description
	return s.repo.Update(current)
}

func (s *RoleService) Delete(id uint) (bool, error) {
	return s.repo.Delete(map[string]interface{}{"id": id})
}
