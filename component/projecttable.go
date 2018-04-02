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
    height="auto"
    style="width: 100%"
    :row-class-name="TableRowClassName"
    @current-change="SetSelectedProject"
    @row-dblclick="SelectRow"
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

	SelectedProject *project.Project   `js:"selected_project"`
	Projects        []*project.Project `js:"projects"`

	VM *hvue.VM `js:"VM"`
}

func NewProjectTableCompModel(vm *hvue.VM) *ProjectTableCompModel {
	ptm := &ProjectTableCompModel{Object: model.O()}
	ptm.Projects = nil
	ptm.SelectedProject = nil
	ptm.VM = vm
	return ptm
}

func (ptm *ProjectTableCompModel) TableRowClassName(rowInfo *js.Object) string {
	p := &project.Project{Object: rowInfo.Get("row")}
	var res string
	switch p.Status {
	case "WiP":
		res = "warning-row"
	case "Done":
		res = "success-row"
	default:
		res = ""
	}
	return res
}

func (ptm *ProjectTableCompModel) SetSelectedProject(np *project.Project) {
	//ptm = &ProjectTableCompModel{Object: vm.Object}
	ptm.SelectedProject = np
	ptm.VM.Emit("update:selected_project", np)
}

func (ptm *ProjectTableCompModel) SelectRow(vm *hvue.VM, prj *project.Project, event *js.Object) {
	vm.Emit("selected_project", prj)
}

func NewProjectTableComp() {
	hvue.NewComponent("project-table",
		hvue.Props("selected_project", "projects"),
		hvue.Template(ProjectTableTemplate),
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			return NewProjectTableCompModel(vm)
		}),
		hvue.MethodsOf(NewProjectTableCompModel(nil)),
		hvue.Computed("nbProjects", func(vm *hvue.VM) interface{} {
			ptm := &ProjectTableCompModel{Object: vm.Object}
			return len(ptm.Projects)
		}),
	)
}
