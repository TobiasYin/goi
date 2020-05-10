package goi

import "syscall/js"

type Event struct {
	// TODO js callback Event
}

type EventCallBack func(e Event)

func createEventWithJsValue(v js.Value) Event {
	// TODO 解开js event
	return Event{}
}

func WrapEventCallBack(cb EventCallBack) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer panicCatch()
		var e Event
		if len(args) > 0 {
			e = createEventWithJsValue(args[0])
		} else {
			e = Event{}
		}
		cb(e)
		return 0
	})
}
