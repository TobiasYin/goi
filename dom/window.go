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

