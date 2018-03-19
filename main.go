package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
)

//go:generate bash ./makejs.sh

func main() {
	go pressAButton()
}

/////////////////////////////////////////////////////////////////////////////////////

type DemoElement struct {
	*js.Object

	Mark    string   `js:"mark"`
	Visible bool     `js:"visible"`
	VM      *hvue.VM `js:"vm"`
}

func NewDemoElement() *DemoElement {
	de := &DemoElement{Object: js.Global.Get("Object").New()}
	de.Mark = "troulala"
	de.Visible = false
	return de
}

func (de *DemoElement) HandleClose(done *js.Object) {
	de.Mark = "itou"
	done.Invoke()
}

func pressAButton() {
	de := NewDemoElement()

	hvue.NewVM(
		hvue.El("#app"),
		hvue.DataS(de),
		hvue.MethodsOf(de),
		hvue.Method("handleClose", func() {
			de.Mark = "itourlilou"
			de.Visible = false
		}),
	)

	js.Global.Set("de", de)
}
