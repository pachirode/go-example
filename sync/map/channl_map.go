package main

type ChannelMap struct {
	cmd chan command
	m   map[string]int
}

type command struct {
	action string // get, set, delete
	key    string
	value  int
	result chan<- result
}

type result struct {
	value int
	ok    bool
}

func NewChannelMap() *ChannelMap {
	sm := &ChannelMap{
		cmd: make(chan command),
		m:   make(map[string]int),
	}

	go sm.run()
	return sm
}

func (m *ChannelMap) run() {
	for cmd := range m.cmd {
		switch cmd.action {
		case "get":
			value, ok := m.m[cmd.key]
			cmd.result <- result{value, ok}
		case "set":
			m.m[cmd.key] = cmd.value
		case "delete":
			delete(m.m, cmd.key)
		}
	}
}

func (m *ChannelMap) Set(key string, value int) {
	m.cmd <- command{action: "set", key: key, value: value}
}

func (m *ChannelMap) Get(key string) (int, bool) {
	res := make(chan result)
	m.cmd <- command{action: "get", key: key, result: res}
	r := <-res
	return r.value, r.ok
}
