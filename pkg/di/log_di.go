package di

import (
	"github.com/chocogem/bigquery-golang/pkg/log"
	"github.com/google/wire"
)

var LogSet = wire.NewSet(
	log.NewImplementLogrusLogger, 
	log.NewLogrusLogger,
	wire.Bind(new(log.Logger),new(*log.LogrusImplement)), 
	
)
