// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateTableView(appData *AppData) *widget.Table {
	table := widget.NewTable(
		// Provide the size of the table
		func() (int, int) {
			log.Printf("NodeList entries %d\n", len(appData.nodeList))
			if len(appData.nodeList) > appData.currentListIndex {
				len := len(appData.nodeList[appData.currentListIndex].Data)
				if len > 0 {
					log.Printf("NodeList chunk len %d\n", len)
					return len/16 + 1, 16
				}
			}
			return 0, 0
		},

		// Provide the content template
		func() fyne.CanvasObject {
			return widget.NewLabel("AA")
		},

		// Provide the content for a specific cell
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if appData.currentListIndex < len(appData.nodeList) {
				idx := i.Row*16 + i.Col
				if idx < len(appData.nodeList[appData.currentListIndex].Data) {
					o.(*widget.Label).SetText(fmt.Sprintf("%02X",
						appData.nodeList[appData.currentListIndex].Data[idx]))
					return
				}
			}
			// Since we always have 16 columns, we must fill the empty ones
			o.(*widget.Label).SetText("")
		},
	)

	return table
}
