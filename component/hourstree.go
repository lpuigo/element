package component

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/model"
	"github.com/lpuig/element/model/node"
)

const HoursTreeTemplate = `
<el-tree
    :data="nodes"
    :props="nodeProps"
>
    <span class="custom-tree-node" slot-scope="{ node, data }">
        <span class="custom-node-name">{{ node.label }}</span>
        <hours-row :hours="data.hours" :hmax="data.maxHour"></hours-row>
    </span>
</el-tree>
`

type HoursTreeCompModel struct {
	*js.Object

	Nodes     []*node.HoursNode `js:"nodes"`
	NodeProps js.M         `js:"nodeProps"`

	VM *hvue.VM `js:"VM"`
}

func NewHoursTreeCompModel(vm *hvue.VM) *HoursTreeCompModel {
	htcm := &HoursTreeCompModel{Object: model.O()}
	htcm.NodeProps = js.M{
		"children": "children",
		"label":    "label",
	}

	htcm.Nodes = []*node.HoursNode{}
	htcm.VM = vm
	return htcm
}

func RegisterHoursTreeComp() {
	//RegisterHoursRowComp()

	hvue.NewComponent("hours-tree",
		hvue.Props("nodes"),
		hvue.Template(HoursTreeTemplate),
		hvue.Component("hours-row", DefineHoursRowComp()...),
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			return NewHoursTreeCompModel(vm)
		}),
		hvue.MethodsOf(&HoursTreeCompModel{}),
	)
}
