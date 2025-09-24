package main

type Message string

func NewMessage() Message {
	return Message("msg")
}

type Greet struct {
	Message Message
}

func NewGreet(msg Message) Greet {
	return Greet{
		Message: msg,
	}
}

func (g Greet) GetMessage() Message {
	return g.Message
}

type Event struct {
	Greet Greet
}

func NewEvent(g Greet) Event {
	return Event{
		Greet: g,
	}
}

func (e Event) Start() {
	println(e.Greet.GetMessage())
}

func main() {
	message := NewMessage()
	greeter := NewGreet(message)
	event := NewEvent(greeter)

	// 使用 wire
	event = InitializeEvent()
	event.Start()
}
