package adapter

import (
	domain "github.com/chocogem/bigquery-golang/pkg/domain"
	repository "github.com/chocogem/bigquery-golang/pkg/repository"
	usecase "github.com/chocogem/bigquery-golang/pkg/usecase"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(cr repository.UserRepository) usecase.UserUseCase {
	return &userUseCase{
		userRepo: cr,
	}
}

func (c *userUseCase) FindAll() ([]domain.User, error) {
	users, err := c.userRepo.FindAll()
	return users, err
}
