package goel

import (
	"github.com/huckridgesw/hvue"
	"github.com/gopherjs/gopherjs/js"
)

var MessageDuration int = 3000

func messageString(vm *hvue.VM, msgtype, msg string, close bool) {
	vm.Call("$message", js.M{
		"showClose" : close,
		"message" : msg,
		"type" : msgtype,
		"duration" : MessageDuration,
	})
}

func MessageInfoStr(vm *hvue.VM, msg string, close bool) {
	messageString(vm, "info", msg, close)
}

func MessageSuccesStr(vm *hvue.VM, msg string, close bool) {
	messageString(vm, "success", msg, close)
}

func MessageWarningStr(vm *hvue.VM, msg string, close bool) {
	messageString(vm, "warning", msg, close)
}

func MessageErrorStr(vm *hvue.VM, msg string, close bool) {
	messageString(vm, "error", msg, close)
}

