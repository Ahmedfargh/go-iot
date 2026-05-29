package crud

import (
	"errors"
	interfaces "go-dashboard/internals/Interfaces"
	models "go-dashboard/internals/Models"
)

type DeviceService struct {
	repo interfaces.RepositoryInterface[models.Device]
}

func NewDeviceService(repo interfaces.RepositoryInterface[models.Device]) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) Create(device *models.Device) (*models.Device, error) {
	existing, _ := s.repo.FindByWhere(map[string]interface{}{"identifier": device.Identifier})
	if len(existing) > 0 {
		return nil, errors.New("a device with this identifier already exists")
	}
	return s.repo.Create(device)
}

func (s *DeviceService) GetByID(id uint) (*models.Device, error) {
	return s.repo.With("Sector").FindById(id)
}

func (s *DeviceService) GetAll(page, pageSize int) ([]models.Device, int64, error) {
	count, _ := s.repo.Count(map[string]interface{}{})
	results, err := s.repo.With("Sector").Paginate(page, pageSize).FindByWhere(map[string]interface{}{})
	return results, count, err
}

func (s *DeviceService) Update(id uint, updatedData *models.Device) (*models.Device, error) {
	current, err := s.repo.FindById(id)
	if err != nil || current == nil {
		return nil, errors.New("device target not found")
	}

	current.Identifier = updatedData.Identifier
	current.Status = updatedData.Status
	current.Temperature = updatedData.Temperature
	current.Humidity = updatedData.Humidity
	current.SectorID = updatedData.SectorID

	return s.repo.Update(current)
}

func (s *DeviceService) Delete(id uint) (bool, error) {
	return s.repo.Delete(map[string]interface{}{"id": id})
}
