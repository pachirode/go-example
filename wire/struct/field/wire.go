//go:build wireinject

package main

import "github.com/google/wire"

func InitializeMessage(phrase string, code int) Content {
	wire.Build(NewMessage, wire.FieldsOf(new(*Message), "Content"))
	return Content("")
}
