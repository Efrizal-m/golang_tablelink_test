package user

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Repository Interface
type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	Upsert(ctx context.Context, ar User) (User, error)
	DeleteByID(ctx context.Context, id int) error
}

// NewRepository will implement MovieRepository Interface
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetBySpec(ctx context.Context) (res []User, err error) {
	query := r.db

	result := query.Find(&res)

	return res, result.Error
}

func (r *repository) Upsert(ctx context.Context, user User) (res User, err error) {
	result := r.db.Save(&user)
	return user, result.Error
}

func (r *repository) DeleteByID(ctx context.Context, id int) (err error) {
	result := r.db.Delete(&User{ID: id})
	return result.Error
}
