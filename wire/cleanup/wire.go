//go:build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeFile(path string) (*App, func(), error) {
	wire.Build(OpenFile, Clean, wire.Struct(new(App), "*"))
	return &App{}, nil, nil
}
