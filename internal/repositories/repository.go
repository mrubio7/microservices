package repositories

import (
	"gorm.io/gorm"
)

type QueryModifier func(*gorm.DB) *gorm.DB

type GenericRepository[T any] struct {
	db *gorm.DB
}

func NewGenericRepository[T any](database *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{db: database}
}

func (r *GenericRepository[T]) Create(entity *T) (*T, error) {
	err := r.db.Create(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *GenericRepository[T]) Update(entity *T, idField string, idValue interface{}) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(entity).Error
}

func (r *GenericRepository[T]) Delete(field string, value interface{}) error {
	var entity T
	return r.db.Where(field+" = ?", value).Delete(&entity).Error
}

func (r *GenericRepository[T]) GetByID(id int32) (*T, error) {
	var entity T
	err := r.db.Model(&entity).First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericRepository[T]) Get(modifiers ...QueryModifier) (*T, error) {
	var entity T
	query := r.db.Model(&entity)

	for _, modify := range modifiers {
		query = modify(query)
	}

	err := query.First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericRepository[T]) Find(modifiers ...QueryModifier) ([]T, error) {
	var entities []T
	query := r.db.Model(&entities)

	// Aplicar todos los modificadores al objeto query
	for _, modify := range modifiers {
		query = modify(query)
	}

	err := query.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *GenericRepository[T]) GetRawQuery(rawQuery string, args ...interface{}) (*T, error) {
	var entity T
	err := r.db.Raw(rawQuery, args...).Scan(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericRepository[T]) FindRawQuery(rawQuery string, args ...interface{}) ([]T, error) {
	var entities []T
	err := r.db.Raw(rawQuery, args...).Scan(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func Where(condition string, values ...interface{}) QueryModifier {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(condition, values...)
	}
}

func OrderBy(order string) QueryModifier {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func Preload(field string) QueryModifier {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(field)
	}
}

func Limit(limit int) QueryModifier {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}
