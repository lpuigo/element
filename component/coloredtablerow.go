package component

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lpuig/element/model"
	"github.com/huckridgesw/hvue"
	"strconv"
	"github.com/lpuig/element/model/colormap"
)

const HoursRowTemplate = `
<div class="custom-hours-row">
    <div class="hours-table">
        <div v-for="h in hours" class="hours-cell" :style="ColorBackground(h)">{{h | FormatHour}}</div>
    </div>
</div>
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
		hvue.Filter("FormatHour", func(vm *hvue.VM, value *js.Object, args ...*js.Object) interface{} {
			h := value.Float()
			return strconv.FormatFloat(h, 'f', 1, 64)
		}),
			hvue.Method("ColorBackground", func(h float64) string {
				c1 := colormap.NewColorFromString("#ff4949")
				c2 := colormap.NewColorFromString("#F7BA2A")
				c3 := colormap.NewColorFromString("#13CE66")

				cm := colormap.NewColorMap()
				cm.Add(0, c1)
				cm.Add(4, c2)
				cm.Add(8, c3)

				return "background-color: " + cm.ColorAt(h).String()
			}),
	)
}
