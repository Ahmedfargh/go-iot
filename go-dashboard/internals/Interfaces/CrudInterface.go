package interfaces

type ServiceInterface[T any] interface {
	Create(model *T) (*T, error)
	GetByID(id uint) (*T, error)
	GetAll(page, pageSize int) ([]T, int64, error)
	Update(id uint, model *T) (*T, error)
	Delete(id uint) (bool, error)
}
