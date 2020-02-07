package router

import (
	"strings"
	"syscall/js"
)

type hashDriver struct {}

func (hd *hashDriver) currentPath() string {
	hashVal := js.Global().Get("location").Get("hash").String()
	return strings.TrimPrefix(hashVal, "#")
}

func (hd *hashDriver) subscribeToChanges(callback func()) (close func()) {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback()
		return nil
	})

	js.Global().Call("addEventListener", "hashchange", f, false)

	return func() {
		js.Global().Call("removeEventListener", "hashchange", f)
		f.Release()
	}
}

func (hd *hashDriver) gotoPath(path string) {
	js.Global().Get("location").Set("hash", path)
}