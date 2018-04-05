package component

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lpuig/element/model"
	"github.com/huckridgesw/hvue"
)

const HoursRowTemplate = `
<table>
    <tr>
        <td v-for="h in hours">{{h}}</td>
    </tr>
</table>
`

type HoursRowModel struct {
	*js.Object

	Hours []float64 `js:"hours"`
}

func NewHoursRowModel() *HoursRowModel {
	res := &HoursRowModel{Object:model.O()}
	res.Hours = []float64{}

	return res
}

func NewColoredTableRowComp() {
	hvue.NewComponent("hours-row",
		hvue.Props("hours"),
		hvue.Template(HoursRowTemplate),
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			return NewHoursRowModel()
		}),
		hvue.MethodsOf(&HoursRowModel{}),
	)
}
