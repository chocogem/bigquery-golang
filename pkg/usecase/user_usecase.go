package usecase

import "github.com/chocogem/bigquery-golang/pkg/domain"

type UserUseCase interface {
	FindAll() ([]domain.User, error)
}
