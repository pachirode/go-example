package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/looplab/fsm"
)

func base() {
	fsmCase := fsm.NewFSM(
		"closed",
		// 如果名字和状态同名，优先当中状态来处理
		fsm.Events{
			{"open", []string{"closed"}, "open"},
			{"close", []string{"open"}, "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsmCase.Current())

	err := fsmCase.Event(context.Background(), "open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsmCase.Current())

	err = fsmCase.Event(context.Background(), "close")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsmCase.Current())
}

func meta() {
	fsmCase := fsm.NewFSM(
		"idle",
		fsm.Events{
			{"produce", []string{"idle"}, "idle"},
			{"consume", []string{"idle"}, "idle"},
			{"remove", []string{"idle"}, "idle"},
		},
		fsm.Callbacks{
			"product": func(_ context.Context, event *fsm.Event) {
				data := "test"
				event.FSM.SetMetadata("message", data)
				fmt.Println("product data: %s\n", data)
			},
			"consume": func(_ context.Context, event *fsm.Event) {
				data, ok := event.FSM.Metadata("message")
				if ok {
					fmt.Println("consume data: %s\n", data)
				}
			},
			"remove": func(_ context.Context, event *fsm.Event) {
				event.FSM.DeleteMetadata("message")
				if _, ok := event.FSM.Metadata("message"); !ok {
					fmt.Println("removed data")
				}
			},
		},
	)

	fmt.Printf("current state: %s\n", fsmCase.Current())

	err := fsmCase.Event(context.Background(), "produce")
	if err != nil {
		fmt.Printf("produce err: %s\n", err)
	}

	fmt.Printf("current state: %s\n", fsmCase.Current())
}

func async() {
	fsmCase := fsm.NewFSM(
		"start",
		fsm.Events{
			{"run", []string{"start"}, "end"},
		},
		fsm.Callbacks{
			"leave_start": func(_ context.Context, event *fsm.Event) {
				event.Async() //标记为异步，触发事件不会进行状态转换，状态转换需要手动调用
			},
		},
	)

	// 不再会触发状态转换
	err := fsmCase.Event(context.Background(), "run")

	var asyncError fsm.AsyncError
	ok := errors.As(err, &asyncError)
	if !ok {
		panic(fmt.Sprintf("not AsyncError"))
	}

	// 手动进行状态转换
	if err = fsmCase.Transition(); err != nil {
		panic(fmt.Sprintf("err when transitioning"))
	}

	fmt.Printf(fsmCase.Current())
}

func main() {
	//base()
	meta()
}
