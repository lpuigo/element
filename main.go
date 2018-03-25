package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"strconv"
	"github.com/lpuig/element/model/project"
	"github.com/lpuig/element/model"
	"github.com/lpuig/element/component"
)

//go:generate bash ./makejs.sh

func main() {
	go pressAButton()
}



/////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////////

type DemoElement struct {
	*js.Object

	Mark    string   `js:"mark"`
	Visible bool     `js:"visible"`
	Projects []*project.Project `js:"projects"`

	VM      *hvue.VM `js:"vm"`
}

func NewDemoElement() *DemoElement {
	de := &DemoElement{Object: model.O()}
	de.Mark = "troulala"
	de.Visible = false
	de.Projects = []*project.Project{}
	status := []string{"Open", "WiP", "Done"}
	for i:= 0; i < 50; i++ {
		de.Projects = append(de.Projects,
			project.NewProject("prj"+strconv.Itoa(i), status[i%3], float64(i)*1.8+1),
		)
	}
	return de
}

func (de *DemoElement) HandleClose(done *js.Object) {
	de.Mark = "itou"
	de.VM.Call("$message", "coucou")
	done.Invoke()
}

func pressAButton() {
	de := NewDemoElement()

	component.NewProjectTableComp()

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

////////////////////////////////////////////////////////////////////////////////////



