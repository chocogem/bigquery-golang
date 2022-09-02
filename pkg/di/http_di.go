package di

import (
	"github.com/chocogem/bigquery-golang/pkg/api"
	"github.com/google/wire"
)

var HTTPSet = wire.NewSet(
	api.NewServerHTTP,
	wire.Struct(new(api.Handlers), "*"),
)
