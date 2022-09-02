package di

import (
	handler "github.com/chocogem/bigquery-golang/pkg/api/handler"
	repoAdapter "github.com/chocogem/bigquery-golang/pkg/repository/adapter"
	usecaseAdapter "github.com/chocogem/bigquery-golang/pkg/usecase/adapter"
	"github.com/google/wire"
)

var FactPostCampaignSet = wire.NewSet(
	repoAdapter.NewUserRepository,
	usecaseAdapter.NewUserUseCase,
	handler.NewUserHandler,
)
