package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget *widget.Entry 
	PreviewWidget *widget.RichText
}

var cnfg config

func main() {
	// Create a fyne app
	a := app.New()

	// Create a window for the app
	w := a.NewWindow("Markdown")

	// Get the user interface
	edit, preview := cnfg.makeUI()

	// Set the content of the window
	w.SetContent(container.NewHSplit(edit, preview))

	// Show window and run app
	w.Resize(fyne.NewSize(500, 500))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func (cnfg *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	cnfg.EditWidget = edit
	cnfg.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}
