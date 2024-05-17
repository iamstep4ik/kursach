package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Моя программка мяу")
	label := widget.NewLabel("Enter smth")
	entry := widget.NewEntry()
	btn := widget.NewButton("Save", func() {
		data := entry.Text
		res := myApp.NewWindow("Result")
		res_label := widget.NewLabel(data)
		res.SetContent(container.NewVBox(res_label))
		res.Show()
	})

	myWindow.SetContent(container.NewVBox(label, entry, btn))

	myWindow.ShowAndRun()
}
