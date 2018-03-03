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

	Visible bool     `js:"visible"`
	VM      *hvue.VM `js:"vm"`
}

func NewDemoElement() *DemoElement {
	de := &DemoElement{Object: js.Global.Get("Object").New()}
	de.Visible = false
	return de
}

func pressAButton() {
	de := NewDemoElement()

	hvue.NewVM(
		hvue.El("#app"),
		hvue.DataS(hvue.NewT(de)),
		hvue.MethodsOf(de),
	)

	js.Global.Set("de", de)
}
