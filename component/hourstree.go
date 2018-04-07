package component

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/model"
)

const HoursTreeTemplate = `
<el-tree
    :data="nodes"
    :props="nodeProps"
    @node-click="HandleNodeClick"
>
    <span class="custom-tree-node" slot-scope="{ node, data }">
        <span class="custom-node-name">{{ node.label }}</span>
        <hours-row :hours="data.hours"></hours-row>
    </span>
</el-tree>
`

type HoursTreeCompModel struct {
	*js.Object

	VM *hvue.VM `js:"VM"`
}

func NewHoursTreeCompModel(vm *hvue.VM) *HoursTreeCompModel {
	htcm := &HoursTreeCompModel{Object: model.O()}

	htcm.VM = vm
	return htcm
}

func NewHoursTreeComp() {
	hvue.NewComponent("hours-tree",
		hvue.Props(""),
		hvue.Template(HoursTreeTemplate),
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			return NewHoursTreeCompModel(vm)
		}),
		hvue.MethodsOf(&HoursTreeCompModel{}),
	)
}
