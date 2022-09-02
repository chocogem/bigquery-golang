package repository

import "github.com/chocogem/bigquery-golang/pkg/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
}
