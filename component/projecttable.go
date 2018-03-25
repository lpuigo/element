package component

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/model"
	"github.com/lpuig/element/model/project"
)

const ProjectTableTemplate = `
<el-table
    :data="projects"
	height="400"
    style="width: 100%"
    :row-class-name="TableRowClassName"
    size="mini"
>
    <el-table-column
        prop="name"
        label="Project Name"
    ></el-table-column>
    <el-table-column
        prop="status"
        label="Status"
    >
    </el-table-column>
    <el-table-column
        prop="workload"
        label="Estimated Work Load"
    >
    </el-table-column>
</el-table>
`

type ProjectTableCompModel struct {
	*js.Object

	Projects []*project.Project `js:"projects"`
}

func (ptm *ProjectTableCompModel) TableRowClassName(rowInfo *js.Object) string {
	p := &project.Project{Object:rowInfo.Get("row")}
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
}

func NewProjectTableCompModel() *ProjectTableCompModel {
	ptm := &ProjectTableCompModel{Object: model.O()}
	ptm.Projects = nil
	return ptm
}

func NewProjectTableComp() {
	hvue.NewComponent("project-table",
		hvue.Props("projects"),
		hvue.Template(ProjectTableTemplate),
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			return NewProjectTableCompModel()
		}),
		hvue.MethodsOf(&ProjectTableCompModel{}),
	)
}
