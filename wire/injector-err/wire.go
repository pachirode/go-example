//go:build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreet, NewMessage)
	return Event{}, nil
}
