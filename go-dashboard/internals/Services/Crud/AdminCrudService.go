package crud

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
)

type AdminService struct {
	repo interfaces.RepositoryInterface[models.Admin]
}

func NewAdminService(repo interfaces.RepositoryInterface[models.Admin]) *AdminService {
	return &AdminService{repo: repo}
}
func (s *AdminService) Create(admin *models.Admin) (*models.Admin, error) {
	existing, _ := s.repo.FindByWhere(map[string]interface{}{"email": admin.Email})
	if len(existing) > 0 {
		return nil, errors.New("this email address is already registered to an admin account")
	}

	existingPhone, _ := s.repo.FindByWhere(map[string]interface{}{"phone": admin.Phone})
	if len(existingPhone) > 0 {
		return nil, errors.New("this phone number is already in use")
	}

	hash := sha256.Sum256([]byte(admin.Password))
	admin.Password = hex.EncodeToString(hash[:])

	return s.repo.Create(admin)
}
func (s *AdminService) GetByID(id uint) (*models.Admin, error) {
	return s.repo.With("Roles", "Permissions").FindById(id)
}
func (s *AdminService) GetAll(page, pageSize int) ([]models.Admin, int64, error) {
	count, _ := s.repo.Count(map[string]interface{}{})
	results, err := s.repo.With("Roles", "Permissions").Paginate(page, pageSize).FindByWhere(map[string]interface{}{})
	return results, count, err
}
func (s *AdminService) Update(id uint, updatedData *models.Admin) (*models.Admin, error) {
	current, err := s.repo.FindById(id)
	if err != nil || current == nil {
		return nil, errors.New("admin profile target not found")
	}
	current.Name = updatedData.Name
	current.Email = updatedData.Email
	current.Phone = updatedData.Phone
	fmt.Println(current.Avatar)
	if updatedData.Avatar != "" {
		current.Avatar = updatedData.Avatar
	}
	return s.repo.Update(current)
}

func (s *AdminService) Delete(id uint) (bool, error) {
	return s.repo.Delete(map[string]interface{}{"id": id})
}
