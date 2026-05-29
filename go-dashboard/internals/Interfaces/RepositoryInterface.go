package interfaces

type RepositoryInterface[T any] interface {
	With(preloads ...string) RepositoryInterface[T]
	Create(model *T) (*T, error)
	Update(model *T) (*T, error)
	FindById(id uint) (*T, error)
	FindByWhere(conditions map[string]interface{}) ([]T, error)
	Delete(conditions map[string]interface{}) (bool, error)
	ForeDelete(conditions map[string]interface{}) (bool, error)
	Restore(conditions map[string]interface{}) (bool, error)
	Paginate(page, pageSize int) RepositoryInterface[T]
	Count(conditions map[string]interface{}) (int64, error)
}
