package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"strconv"
)

//go:generate bash ./makejs.sh

func main() {
	go pressAButton()
}



/////////////////////////////////////////////////////////////////////////////////////

func O() *js.Object {
	return js.Global.Get("Object").New()
}

/////////////////////////////////////////////////////////////////////////////////////

type DemoElement struct {
	*js.Object

	Mark    string   `js:"mark"`
	Visible bool     `js:"visible"`
	Projects []*Project `js:"projects"`

	VM      *hvue.VM `js:"vm"`
}

func NewDemoElement() *DemoElement {
	de := &DemoElement{Object: O()}
	de.Mark = "troulala"
	de.Visible = false
	de.Projects = []*Project{}
	status := []string{"Open", "WiP", "Done"}
	for i:= 0; i < 150; i++ {
		de.Projects = append(de.Projects,
			NewProject("prj"+strconv.Itoa(i), status[i%3], float64(i)*1.8+1),
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

	hvue.NewVM(
		hvue.El("#app"),
		hvue.DataS(de),
		hvue.MethodsOf(de),
		hvue.Method("handleClose", func() {
			de.Mark = "itourlilou"
			de.Visible = false
		}),
		hvue.Method("tableRowClassName", func (rowInfo *js.Object) string {
			p := &Project{Object:rowInfo.Get("row")}
			var res string
			switch p.Status {
			case "WiP":
				res = "warning-row"
			case "Done":
				res = "success-row"
			default:
				res = ""
			}
			println("retrieved project :", p.Object, res)
			return res
		}),
	)

	js.Global.Set("de", de)
}

////////////////////////////////////////////////////////////////////////////////////

type Project struct {
	*js.Object

	Name     string  `js:"name"`
	Workload float64 `js:"workload"`
	Status   string  `js:"status"`
}

func NewProject(name, status string, wl float64) *Project {
	p := &Project{Object:O()}
	p.Name = name
	p.Workload = wl
	p.Status = status

	return p
}


