package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/component"
	"github.com/lpuig/element/model"
	"github.com/lpuig/element/model/project"
	"strconv"
)

//go:generate bash ./makejs.sh

func main() {
	go pressAButton()
}

/////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////////

type DemoElement struct {
	*js.Object

	EditedProject *project.Project   `js:"editedProject"`
	Projects      []*project.Project `js:"projects"`

	VM *hvue.VM `js:"vm"`
}

func NewDemoElement() *DemoElement {
	de := &DemoElement{Object: model.O()}
	de.EditedProject = nil
	de.Projects = []*project.Project{}
	status := []string{"Open", "WiP", "Done"}
	for i := 0; i < 5; i++ {
		de.Projects = append(de.Projects,
			project.NewProject("prj"+strconv.Itoa(i), status[i%3], float64(i)*1.8+1),
		)
	}
	return de
}

func (de *DemoElement) EditProject(prj *project.Project) {
	de.VM.Refs("ProjectEdit").Call("Show", prj)
}

func (de *DemoElement) AddProject(prj *project.Project) {
	de.Projects = append(de.Projects, prj)
}

func (de *DemoElement) RemoveProject(prj *project.Project) {
	for i, p := range de.Projects {
		if p.Object == prj.Object {
			de.Object.Get("projects").Call("splice", i, 1)
		}
	}
}


func pressAButton() {
	de := NewDemoElement()

	component.NewProjectTableComp()
	component.NewProjectEditComp()

	hvue.NewVM(
		hvue.El("#app"),
		hvue.DataS(de),
		hvue.MethodsOf(de),
	)

	js.Global.Set("de", de)
}

////////////////////////////////////////////////////////////////////////////////////
