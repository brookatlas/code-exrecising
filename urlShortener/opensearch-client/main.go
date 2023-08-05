package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/fstanis/screenresolution"
)

var a fyne.App
var window fyne.Window
var windowTitle string

func main() {
	resolution := screenresolution.GetPrimary()
	if resolution == nil {
		fmt.Println("failed to get screen resolution")
		os.Exit(1)
	}
	windowTitle = "Opensearch Client"
	a = app.New()
	window = a.NewWindow(windowTitle)

	windowSize := fyne.Size{
		float32(resolution.Width),
		float32(resolution.Height),
	}
	window.Resize(windowSize)

	// toolbarItems := []widget.ToolbarItem{
	// 	widget.NewToolbarAction(widget.NewLabel("file"), func() {
	// 		fmt.Println("file action!")
	// 	}),
	// }

	toolbarItems := []fyne.CanvasObject{
		widget.NewButton(
			"connection",
			onConnectionButtonClick,
		),
		widget.NewButton(
			"about",
			onAboutButtonClick,
		),
	}

	toolbarGridContainer := container.NewHBox(toolbarItems...)
	mainContainer := container.NewBorder(
		toolbarGridContainer,
		nil,
		nil,
		nil,
	)

	window.SetContent(mainContainer)

	window.ShowAndRun()
}

func onConnectionButtonClick() {
	fmt.Println("connection clicked!")
	dialogContent := container.NewVBox(
		widget.NewLabel("hello"),
	)
	dialog.ShowCustom(
		"new connection form",
		"dismiss",
		dialogContent,
		window,
	)
	// dialog.ShowCustom(
	// 	"new connection",
	// 	"",
	// 	container.NewVBox(),
	// 	window,
	// )
}

func onAboutButtonClick() {
	fmt.Println("about was clicked!")
}
