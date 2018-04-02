package component

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/model"
	"github.com/lpuig/element/model/project"
)

const ProjectEditTemplate string = `
<el-dialog :visible.sync="visible" width="60%">
    <!--<span slot="title" class="dialog-header">-->
    <span slot="title">
        <h2 v-if="edited_project" style="margin: 0 0">Edit Project: <span style="color: teal">{{edited_project.name}}</span></h2>
    </span>


    <el-form v-if="currentProject"
        :model="currentProject" 
        label-width="150px" 
        size="mini"
    >
        <el-form-item label="Project Name">
            <el-input v-model="currentProject.name"></el-input>
        </el-form-item>
        <el-form-item label="Project Status">
            <el-select v-model="currentProject.status" placeholder="project status">
                <el-option label="Open" value="Open"></el-option>
                <el-option label="On Going" value="WiP"></el-option>
                <el-option label="Finished" value="Done"></el-option>
            </el-select>
        </el-form-item>
        <el-form-item label="Project WorkLoad">
            <el-input-number v-model="currentProject.workload"
                             min=0
                             step=0.5
            ></el-input-number>
        </el-form-item>
    </el-form>

    <span slot="footer" class="dialog-footer">
        <el-popover 
                v-if="!isNewProject"
                ref="popoverdelete"
                placement="top"
                width="160"
                v-model="showconfirmdelete"
                :disable="!visible"
        >
            <p>Are you sure to delete this project ?</p>
            <div style="text-align: right; margin: 0">
                <el-button size="mini" type="text" @click="showconfirmdelete = false">nope</el-button>
                <el-button size="mini" type="primary" @click="RemoveProject">yup</el-button>
            </div>
        </el-popover>
    
        <el-button v-if="!isNewProject" type="danger" icon="el-icon-delete" circle v-popover:popoverdelete></el-button>
        <el-button v-if="!isNewProject" @click="Duplicate">Duplicate</el-button>
        <el-button @click="visible = false">Cancel</el-button>
        <el-button v-if="!isNewProject" type="primary" @click="ConfirmChange">Confirm Change</el-button>
        <el-button v-if="isNewProject" type="primary" @click="NewProject">Add New</el-button>
    </span>
</el-dialog>
`

type ProjectEditCompModel struct {
	*js.Object

	EditedProject     *project.Project `js:"edited_project"`
	CurrentProject    *project.Project `js:"currentProject"`

	Visible           bool             `js:"visible"`
	IsNewProject      bool             `js:"isNewProject"`
	ShowConfirmDelete bool             `js:"showconfirmdelete"`

	VM *hvue.VM `js:"VM"`
}

func NewProjectEditCompModel(vm *hvue.VM) *ProjectEditCompModel {
	pecm := &ProjectEditCompModel{Object: model.O()}
	pecm.EditedProject = project.NewProject("", "", 0)
	pecm.CurrentProject = project.NewProject("", "", 0)
	pecm.Visible = false
	pecm.IsNewProject = false
	pecm.ShowConfirmDelete = false
	pecm.VM = vm
	return pecm
}

func (pecm *ProjectEditCompModel) Show(p *project.Project) {
	pecm.EditedProject = p
	pecm.CurrentProject = project.NewProject("", "", 0)
	pecm.CurrentProject.Copy(pecm.EditedProject)
	pecm.IsNewProject = false
	pecm.ShowConfirmDelete = false
	pecm.Visible = true
}

func (pecm *ProjectEditCompModel) ConfirmChange() {
	pecm.EditedProject.Copy(pecm.CurrentProject)
	pecm.VM.Emit("update:edited_project", pecm.EditedProject)
	pecm.Visible = false
}

func (pecm *ProjectEditCompModel) Duplicate() {
	pecm.EditedProject = pecm.CurrentProject
	pecm.CurrentProject.Name += " (Copy)"
	pecm.IsNewProject = true
}

func (pecm *ProjectEditCompModel) NewProject() {
	pecm.VM.Emit("newproject", pecm.CurrentProject)
	pecm.Visible = false
}

func (pecm *ProjectEditCompModel) RemoveProject() {
	pecm.VM.Emit("removeproject", pecm.EditedProject)
	pecm.ShowConfirmDelete = false
	pecm.Visible = false
}

func NewProjectEditComp() {
	hvue.NewComponent("project-edit",
		hvue.Props("edited_project"),
		hvue.Template(ProjectEditTemplate),
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			return NewProjectEditCompModel(vm)
		}),
		hvue.MethodsOf(&ProjectEditCompModel{}),
	)
}
