//go:build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeEvent() Event {
	var providerSet wire.ProviderSet = wire.NewSet(NewMessage, NewGreet, NewMessage)
	wire.Build(providerSet)
	return Event{}
}
