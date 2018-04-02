package project

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lpuig/element/model"
)

type Project struct {
	*js.Object

	Name     string  `js:"name"`
	Workload float64 `js:"workload"`
	Status   string  `js:"status"`
}

func NewProject(name, status string, wl float64) *Project {
	p := &Project{Object: model.O()}
	p.Name = name
	p.Workload = wl
	p.Status = status

	return p
}

func ProjectFromJS(o *js.Object) *Project {
	return &Project{Object: o}
}

func (p *Project) Copy(refProj *Project) {
	p.Name = refProj.Name
	p.Workload = refProj.Workload
	p.Status = refProj.Status
}
