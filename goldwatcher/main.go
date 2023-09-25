package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App        *fyne.App
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	MainWindow fyne.Window
}

var myApp Config

func main() {
	// create a fyne app
	a := app.NewWithID("com.ritesh.preferences")
	myApp.App = a

	// create our loggers
	myApp.InfoLog = log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open connection to the database

	// create a database repository

	// create and resize the window
	myApp.MainWindow = a.NewWindow("Gold Wacher")
	myApp.MainWindow.Resize(fyne.NewSize(300, 200))
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// show and run the app
	myApp.MainWindow.ShowAndRun()
}
