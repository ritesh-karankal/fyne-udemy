package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cnfg config

func main() {
	// Create a fyne app
	a := app.New()

	// a.Settings().SetTheme(&myTheme{})

	// Create a window for the app
	w := a.NewWindow("Markdown")

	// Get the user interface
	edit, preview := cnfg.makeUI()

	cnfg.createMenuItems(w)

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

func (cnfg *config) createMenuItems(win fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", cnfg.openFunc(win))
	saveMenuItem := fyne.NewMenuItem("Save", cnfg.saveFunc(win))

	cnfg.SaveMenuItem = saveMenuItem
	cnfg.SaveMenuItem.Disabled = true

	saveAsMenuItem := fyne.NewMenuItem("Save As...", cnfg.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	menu := fyne.NewMainMenu(fileMenu)
	win.SetMainMenu(menu)
}

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveAsDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				// user cancelled
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with .md extension", win)
				return
			}

			// save file
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, win)

		saveAsDialog.SetFileName("untitled.md")
		saveAsDialog.SetFilter(filter)
		saveAsDialog.Show()
	}
}

func (app *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				// user cancelled
				return
			}

			defer read.Close()

			data, err := ioutil.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			app.EditWidget.SetText(string(data))

			app.CurrentFile = read.URI()
			win.SetTitle(win.Title() + " - " + read.URI().Name())

			app.SaveMenuItem.Disabled = false

		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *config) saveFunc(win fyne.Window) func() {
	return func() {
		if app.CurrentFile != nil {
			write, err := storage.Writer(app.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(app.EditWidget.Text))
			defer write.Close()

		}
	}
}
