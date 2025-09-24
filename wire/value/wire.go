//go:build wireinject

package main

import "github.com/google/wire"

func InitializeMessage(phrase string, code int) Message {
	wire.Build(wire.Value(Message{
		Message: "msg",
	}))
	return Message{}
}
