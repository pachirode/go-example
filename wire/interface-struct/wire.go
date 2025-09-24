//go:build wireinject

package main

import "github.com/google/wire"

func WireRunStore(msg *Message) error {
	wire.Build(SaveMessage, New, wire.Bind(new(Store), new(*store)))
	return nil
}
