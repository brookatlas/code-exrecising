package main

import (
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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

	windowSize := fyne.NewSize(
		float32(resolution.Width),
		float32(resolution.Height),
	)
	window.Resize(windowSize)
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
	connectionListWidget := widget.NewList(
		func() int {
			return 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("example connection")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {},
	)

	connectionListBackgroundContainer := container.NewMax(
		canvas.NewRectangle(color.RGBA{80, 80, 80, 255}),
		connectionListWidget,
	)

	mainContainer := container.NewBorder(
		toolbarGridContainer,
		nil,
		connectionListBackgroundContainer,
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

	dialogContainer := container.NewMax(
		widget.NewLabel(
			"Opensearch Client by barak atias. \n version 0.1",
		),
	)
	aboutDialog := dialog.NewCustom(
		"About",
		"close",
		dialogContainer,
		window,
	)

	aboutDialog.Show()
}
