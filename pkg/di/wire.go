//go:build wireinject
// +build wireinject

package di

import (
	"github.com/chocogem/bigquery-golang/pkg/api"
	"github.com/chocogem/bigquery-golang/pkg/config"
	"github.com/google/wire"
)
func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(FactPostCampaignSet, HTTPSet)

	return &api.ServerHTTP{}, nil
}
