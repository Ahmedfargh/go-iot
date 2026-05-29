package repository

import (
	"errors"
	interfaces "go-dashboard/internals/Interfaces"

	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db       *gorm.DB
	preloads []string
	page     int
	pageSize int
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) With(preloads ...string) interfaces.RepositoryInterface[T] {
	return &GormRepository[T]{
		db:       r.db,
		preloads: append(r.preloads, preloads...),
		page:     r.page,
		pageSize: r.pageSize,
	}
}

func (r *GormRepository[T]) Paginate(page, pageSize int) interfaces.RepositoryInterface[T] {
	return &GormRepository[T]{
		db:       r.db,
		preloads: r.preloads,
		page:     page,
		pageSize: pageSize,
	}
}

func (r *GormRepository[T]) applyPreloads(db *gorm.DB) *gorm.DB {
	for _, preload := range r.preloads {
		db = db.Preload(preload)
	}
	if r.page > 0 && r.pageSize > 0 {
		offset := (r.page - 1) * r.pageSize
		db = db.Offset(offset).Limit(r.pageSize)
	}
	return db
}

func (r *GormRepository[T]) Create(model *T) (*T, error) {
	err := r.db.Create(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (r *GormRepository[T]) Update(model *T) (*T, error) {
	err := r.db.Save(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (r *GormRepository[T]) FindById(id uint) (*T, error) {
	var entity T
	db := r.applyPreloads(r.db)
	err := db.First(&entity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}
func (r *GormRepository[T]) FindByWhere(conditions map[string]interface{}) ([]T, error) {
	var results []T
	db := r.applyPreloads(r.db)
	err := db.Where(conditions).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r *GormRepository[T]) Delete(conditions map[string]interface{}) (bool, error) {
	var entity T
	result := r.db.Where(conditions).Delete(&entity)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
func (r *GormRepository[T]) ForeDelete(conditions map[string]interface{}) (bool, error) {
	var entity T
	result := r.db.Unscoped().Where(conditions).Delete(&entity)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
func (r *GormRepository[T]) Restore(conditions map[string]interface{}) (bool, error) {
	var entity T
	result := r.db.Unscoped().Where(conditions).Model(&entity).Update("deleted_at", nil)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
func (r *GormRepository[T]) Count(conditions map[string]interface{}) (int64, error) {
	var count int64
	var entity T
	err := r.db.Model(&entity).Where(conditions).Count(&count).Error
	return count, err
}
