package main

import (
	"errors"
	"fmt"
	"time"
)

type Content string

type Message struct {
	Content Content
}

func NewMessage(content string) *Message {
	return &Message{
		Content: Content(content),
	}
}

type Greet struct {
	Message Message
}

func NewGreet(m Message) Greet {
	return Greet{Message: m}
}

func (g Greet) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greet
}

func NewEvent(g Greet) (Event, error) {
	if time.Now().Unix()%2 == 0 {
		return Event{}, errors.New("new event error")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
