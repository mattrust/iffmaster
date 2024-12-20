// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/mattrust/iffmaster/internal/chunks"
)

type ListEntry struct {
	label            string
	description      string
	structure        chunks.StructResult
	*chunks.IFFChunk // Embedding the IFFChunk struct
}

func NewListView(appData *AppData) *widget.List {
	list := widget.NewList(

		// The number of items in the list
		func() int {
			return len(appData.nodeList)
		},

		// The function to create the widget for each item
		func() fyne.CanvasObject {
			return widget.NewLabel("WWWWWWWW")
		},

		// The function to populate the widget with the data for each item
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			entry := appData.nodeList[i]
			obj.(*widget.Label).SetText(entry.label)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		appData.chunkInfo.SetText(appData.nodeList[id].description)
		appData.currentListIndex = id
		appData.hexTableView.Refresh()
		appData.isoTableView.Refresh()
		appData.structTableView.Refresh()
	}

	return list
}

func ConvertIFFChunkToListNode(chunk *chunks.IFFChunk) []ListEntry {
	var nodeList []ListEntry

	var traverse func(chunk *chunks.IFFChunk, level int)
	traverse = func(chunk *chunks.IFFChunk, level int) {
		indentation := ""
		for i := 0; i < level; i++ {
			indentation += "."
		}
		nodeList = append(nodeList, ListEntry{
			label: indentation + chunk.ID,
			description: fmt.Sprintf(
				"Type: %s - Desc.: %s - Size: %d",
				chunk.ChType, chunks.GetChunkDescription(chunk.ChType), chunk.Size),
			IFFChunk:  chunk,
			structure: chunks.GetChunkStructure(chunk.ChType, chunk.Data)})
		for _, child := range chunk.Childs {
			traverse(child, level+1)
		}
	}

	traverse(chunk, 0)
	return nodeList
}
