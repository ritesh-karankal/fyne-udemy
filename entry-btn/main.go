package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello!")

	output, entry, btn := makeUI()
	
	w.SetContent(container.NewVBox(output, entry, btn))
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}

func makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello World!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		output.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance
	
	return output, entry, btn
}
