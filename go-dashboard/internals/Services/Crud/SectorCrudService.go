package crud

import (
	"errors"
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
)

type SectorService struct {
	repo interfaces.RepositoryInterface[models.Sector]
}

func NewSectorService(repo interfaces.RepositoryInterface[models.Sector]) *SectorService {
	return &SectorService{repo: repo}
}

func (s *SectorService) Create(sector *models.Sector) (*models.Sector, error) {
	existing, _ := s.repo.FindByWhere(map[string]interface{}{"name": sector.Name})
	if len(existing) > 0 {
		return nil, errors.New("a sector with this name already exists")
	}
	return s.repo.Create(sector)
}

func (s *SectorService) GetByID(id uint) (*models.Sector, error) {
	return s.repo.With("Children", "Parent", "Devices").FindById(id)
}

func (s *SectorService) GetAll(page, pageSize int) ([]models.Sector, int64, error) {
	count, _ := s.repo.Count(map[string]interface{}{})
	results, err := s.repo.With("Parent").Paginate(page, pageSize).FindByWhere(map[string]interface{}{})
	return results, count, err
}

func (s *SectorService) Update(id uint, updatedData *models.Sector) (*models.Sector, error) {
	current, err := s.repo.FindById(id)
	if err != nil || current == nil {
		return nil, errors.New("sector target not found")
	}

	if updatedData.ParentID != nil && *updatedData.ParentID == id {
		return nil, errors.New("a sector cannot be its own parent")
	}

	current.Name = updatedData.Name
	current.Description = updatedData.Description
	current.ParentID = updatedData.ParentID

	return s.repo.Update(current)
}

func (s *SectorService) Delete(id uint) (bool, error) {
	return s.repo.Delete(map[string]interface{}{"id": id})
}
