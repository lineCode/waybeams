package main

import (
	. "display"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func createWindow() (Displayable, error) {
	return GlfwWindow(NewBuilder(), Title("Test Title"), Width(640), Height(480), GlfwFrameRate(10), Children(func(b Builder) {
		Box(b, Id("header"), FlexHeight(1), FlexWidth(1))
		HBox(b, Id("body"), FlexHeight(1), FlexWidth(1), Children(func(b Builder) {
			Box(b, Id("leftNav"), FlexWidth(1), FlexHeight(1))
			Box(b, Id("content"), FlexWidth(3), FlexHeight(1))
		}))
		Box(b, Id("footer"), FlexHeight(1), FlexWidth(1))
	}))
}

func main() {
	win, err := createWindow()
	if err != nil {
		panic(err)
	}
	win.(*GlfwWindowComponent).Loop()
}
