// wire.go
//go:build wireinject
// +build wireinject

package service

import (
	"gvb/config"
	"gvb/core"
	"gvb/dao"

	"github.com/google/wire"
)

var SvcSet = wire.NewSet(
	config.ProvideSet,
	core.ProvideSet,
	dao.ProvideSet,
	ProvideSet,
)

func New(config config.Config) (*Service, error) {
	wire.Build(SvcSet)
	return &Service{}, nil
}
