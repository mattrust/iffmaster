// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

// Package gui provides the GUI for the IFF Master application
package gui

import (
	"bytes"
	"io"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mattrust/iffmaster/internal/chunks"
)

type AppData struct {
	chunks *chunks.IFFChunk

	app          fyne.App
	win          fyne.Window
	topContainer *fyne.Container

	listView *widget.List
	nodeList []ListEntry

	currentListIndex int

	chunkInfo *widget.Label

	hexTableView    *widget.Table
	isoTableView    *widget.Table
	structTableView *widget.Table
}

// OpenGUI layouts the main window and opens it.
// If a filename is given, it reads the file and displays its content.
// The version string is used to display the version of the application.
func OpenGUI(filename string, version string) {
	var appData AppData
	var fileDlg *dialog.FileDialog

	appData.nodeList = make([]ListEntry, 0)

	appData.app = app.NewWithID("github.mattrust.iffmaster")
	appData.win = appData.app.NewWindow("IFF Master")

	toolBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), func() {
			fileDlg = dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil || reader == nil {
					return
				}

				// reset AppData and GUI
				appData.nodeList = make([]ListEntry, 0)
				appData.currentListIndex = 0
				appData.chunkInfo.SetText("")
				appData.listView.UnselectAll()

				appData.topContainer.Refresh()

				// read the file to get its length
				data, err := io.ReadAll(reader)
				if err != nil {
					dialog.ShowError(err, appData.win)
					return
				}

				appData.chunks, err = chunks.ReadIFFFile(bytes.NewReader(data),
					int64(len(data)))
				if err != nil {
					dialog.ShowError(err, appData.win)
					return
				}
				chunks.PrintIffChunk(appData.chunks, 0)

				appData.nodeList = ConvertIFFChunkToListNode(appData.chunks)
				appData.listView.UnselectAll()
				appData.topContainer.Refresh()
			}, appData.win)
			fileDlg.Show()
		}),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			dialog.ShowInformation("About",
				"IFF Master\n"+
					"Version: v"+version+"\n\n"+
					"An Open Source tool for inspecting of IFF(EA-85)-Dateien.\n\n"+
					"License: MIT.\n\n"+
					"Developed by Matthias Rustler and the Open Source community.\n"+
					"Github project page https://github.com/mattrust/iffmaster.", appData.win)
		}),
	)

	appData.listView = NewListView(&appData)
	appData.hexTableView = NewHexTableView(&appData)
	appData.isoTableView = NewIsoTableView(&appData)
	appData.structTableView = NewStructTableView(&appData)

	tabs := container.NewAppTabs(
		container.NewTabItem("Hex", appData.hexTableView),
		container.NewTabItem("ISO8859-1", appData.isoTableView),
		container.NewTabItem("Structure", appData.structTableView))

	appData.chunkInfo = widget.NewLabel("")

	cont1 := container.NewBorder(appData.chunkInfo, nil, nil, nil, tabs)
	appData.topContainer = container.NewBorder(toolBar, nil, appData.listView, nil, cont1)
	appData.win.SetContent(appData.topContainer)

	appData.win.Resize(fyne.NewSize(800, 600))

	readFileName(&appData, filename)

	appData.win.ShowAndRun()
}

// readFileName reads the file with the given filename and displays its content.
// If the filename is empty, it does nothing.
func readFileName(appData *AppData, filename string) {
	log.Print("readFileName: ", filename)
	if filename != "" {
		data, err := os.ReadFile(filename)
		if err != nil {
			dialog.ShowError(err, appData.win)
			return
		}

		appData.chunks, err = chunks.ReadIFFFile(bytes.NewReader(data), int64(len(data)))
		if err != nil {
			dialog.ShowError(err, appData.win)
			return
		}

		appData.nodeList = ConvertIFFChunkToListNode(appData.chunks)
		appData.topContainer.Refresh()
	}
}
