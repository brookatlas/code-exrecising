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
	dialogSize := fyne.NewSize(
		300,
		300,
	)

	connectionNameLabel := widget.NewLabel(
		"connection name:",
	)
	connectionNameEntry := widget.NewEntry()

	hostnameLabel := widget.NewLabel(
		"hostname:",
	)
	hostnameEntry := widget.NewEntry()

	usernameLabel := widget.NewLabel(
		"username:",
	)
	usernameEntry := widget.NewEntry()

	passwordLabel := widget.NewLabel(
		"password:",
	)
	passwordEntry := widget.NewPasswordEntry()

	dialogContent := container.NewVBox(
		container.NewMax(
			container.NewGridWithColumns(
				2,
				connectionNameLabel,
				connectionNameEntry,
			),
		),
		container.NewMax(
			container.NewGridWithColumns(
				2,
				hostnameLabel,
				hostnameEntry,
			),
		),
		container.NewMax(
			container.NewGridWithColumns(
				2,
				usernameLabel,
				usernameEntry,
			),
		),
		container.NewMax(
			container.NewGridWithColumns(
				2,
				passwordLabel,
				passwordEntry,
			),
		),
	)

	dialog := dialog.NewCustom(
		"New Connection",
		"dismiss",
		dialogContent,
		window,
	)
	dialog.Resize(
		dialogSize,
	)

	dialog.Show()
}

func onAboutButtonClick() {
	fmt.Println("about was clicked!")
}
