// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package gui

import (
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

	app fyne.App
	win fyne.Window

	listView *widget.List
	nodeList []ListEntry

	currentListIndex int

	chunkInfo *widget.Label

	hexTableView *widget.Table
	isoTableView *widget.Table
}

func OpenGUI() {
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
				appData.hexTableView.Refresh()
				appData.isoTableView.Refresh()

				appData.chunks, err = chunks.ReadIFFFile(reader)
				if err != nil {
					dialog.ShowError(err, appData.win)
					return
				}
				chunks.PrintIffChunk(appData.chunks, 0)

				appData.nodeList = ConvertIFFChunkToListNode(appData.chunks)
				appData.listView.Refresh()
				appData.hexTableView.Refresh()
				appData.isoTableView.Refresh()
			}, appData.win)
			fileDlg.Show()
		}),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			dialog.ShowInformation("About", "IFF Master\nVersion 1.0", appData.win)
		}),
	)

	appData.listView = NewListView(&appData)
	appData.hexTableView = NewHexTableView(&appData)
	appData.isoTableView = NewIsoTableView(&appData)

	tabs := container.NewAppTabs(
		container.NewTabItem("Hex", appData.hexTableView),
		container.NewTabItem("ISO8859-1", appData.isoTableView),
	)

	appData.chunkInfo = widget.NewLabel("")

	cont1 := container.NewBorder(appData.chunkInfo, nil, nil, nil, tabs)
	appData.win.SetContent(container.NewBorder(toolBar, nil, appData.listView, nil, cont1))

	appData.win.Resize(fyne.NewSize(800, 600))
	appData.win.ShowAndRun()
}
