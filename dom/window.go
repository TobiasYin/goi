package dom

import (
	"math"
	"syscall/js"
)

type Window struct {
	Value
}

func (w Window) ParseFloat(v js.Value) float64 {
	res := w.call(v).Float()
	if math.IsNaN(res) {
		res = 0
	}
	return res
}

func (w Window) GetHash() string {
	return w.Get("location").Get("hash").String()
}

func (w Window) AddEventListener(name string, event js.Func) {
	w.call(name, event)
}

func (w Window) SetHash(hash string) {
	w.Get("location").Set("hash", hash)
}

func (w Window) InnerWidth() int {
	return w.Get("innerWidth").Int()
}

func (w Window) InnerHeight() int {
	return w.Get("innerHeight").Int()
}

func (w Window) OuterWidth() int {
	return w.Get("outerWidth").Int()
}

func (w Window) OuterHeight() int {
	return w.Get("outerHeight").Int()
}
