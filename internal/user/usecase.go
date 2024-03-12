package user

import (
	"context"
)

// UseCase ...
type UseCase interface {
	FindAll(context context.Context) ([]User, error)
	Add(context context.Context, user User) (User, error)
	Update(context context.Context, editedUser User) (User, error)
	Delete(context context.Context, id int) error
}

type useCase struct {
	repo Repository
}

// NewUseCase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewUseCase(repo Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

// FindByTitle ...
func (us *useCase) FindAll(context context.Context) (res []User, err error) {
	res, err = us.repo.GetAll(context)
	return res, err
}

// Add ...
func (us *useCase) Add(context context.Context, newUser User) (res User, err error) {
	res, err = us.repo.Upsert(context, newUser)
	return res, err
}

// Update ...
func (us *useCase) Update(context context.Context, editedUser User) (res User, err error) {
	res, err = us.repo.Upsert(context, editedUser)
	return res, err
}

// Remove ...
func (us *useCase) Delete(context context.Context, id int) (err error) {
	err = us.repo.DeleteByID(context, id)
	return err
}
