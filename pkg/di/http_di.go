package di

import (
	"github.com/chocogem/bigquery-golang/pkg/api"
	"github.com/chocogem/bigquery-golang/pkg/api/middleware"
	"github.com/google/wire"
)

var HTTPSet = wire.NewSet(
	api.NewServerHTTP,
	middleware.NewAuthenticationHandler,
	middleware.NewErrorHandler,
	wire.Struct(new(api.Middlewares), "*"),
	wire.Struct(new(api.Handlers), "*"),
)
