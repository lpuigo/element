package component

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/huckridgesw/hvue"
	"github.com/lpuig/element/model"
	"strconv"
	"github.com/lpuig/element/model/colormap"
)

const HoursRowTemplate = `
<div class="custom-hours-row">
    <div class="hours-table">
        <div 
                v-for="h in hours" 
                class="hours-cell" 
                :style="ColorBackground(h/hmax)"
        >{{h | FormatHour}}</div>
    </div>
</div>
`

type HoursRowModel struct {
	*js.Object

	Hours []float64 `js:"hours"`
	MaxHour float64 `js:"hmax"`
}

func NewHoursRowModel() *HoursRowModel {
	res := &HoursRowModel{Object: model.O()}
	res.Hours = []float64{}

	return res
}

func RegisterHoursRowComp() {
	hvue.NewComponent("hours-row",
		DefineHoursRowComp()...
	)
}

func DefineHoursRowComp() []hvue.ComponentOption {
	return []hvue.ComponentOption{
		hvue.Props("hours", "hmax"),
		hvue.Template(HoursRowTemplate),
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			return NewHoursRowModel()
		}),
		hvue.MethodsOf(&HoursRowModel{}),
		hvue.Filter("FormatHour", func(vm *hvue.VM, value *js.Object, args ...*js.Object) interface{} {
			h := value.Float()
			return strconv.FormatFloat(h, 'f', 0, 64)
		}),
		hvue.Method("ColorBackground", func(h float64) string {
			c1 := colormap.NewColorFromString("#ff4949")
			c2 := colormap.NewColorFromString("#F7BA2A")
			c3 := colormap.NewColorFromString("#13CE66")

			cm := colormap.NewColorMap()
			cm.Add(0, c1)
			cm.Add(0.5, c2)
			cm.Add(1, c3)

			return "background-color: " + cm.ColorAt(h).String()
		}),
	}
}