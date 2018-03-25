package component

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/model/project"
	"github.com/lpuig/element/model"
)

const ProjectTableTemplate = `
<el-table
    :data="projects"
    style="width: 95%"
    :row-class-name="tableRowClassName"
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

func NewProjectTableCompModel() *ProjectTableCompModel {
	ptm := &ProjectTableCompModel{Object:model.O()}
	ptm.Projects = nil
	return ptm
}

func NewProjectTableComp() hvue.ComponentOption {
	return hvue.Component("project-table",
		hvue.Template(ProjectTableTemplate),
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			return NewProjectTableCompModel()
		}),
	)
}
