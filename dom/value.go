package dom

import (
	"fmt"
	"runtime"
	"strings"
	"syscall/js"
)


type Value struct {
	js.Value
}

func (v Value) callWithPC(pc uintptr, args ...interface{}) js.Value {
	name := runtime.FuncForPC(pc).Name()
	index := strings.LastIndex(name, ".")
	if index >= 0 {
		name = name[index+1:]
	}
	name = strings.ToLower(name[0:1]) + name[1:]
	fmt.Printf("Call %v\n", name)
	return v.Call(name, args...)
}

func (v Value) call(args ...interface{}) js.Value {
	pc, _, _, _ := runtime.Caller(1)
	return v.callWithPC(pc, args...)
}

func (v Value) callWithTwoSkip(args ...interface{}) js.Value {
	pc, _, _, _ := runtime.Caller(2)
	return v.callWithPC(pc, args...)
}