//go:build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreet, wire.Struct(new(Message), "Content"))
	return Event{}, nil
}
