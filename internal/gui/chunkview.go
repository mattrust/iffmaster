// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateTableView(appData *AppData) *widget.Table {
	table := widget.NewTable(
		// Provide the size of the table
		func() (int, int) {
			//fmt.Printf("NodeList: %d\n", len(appData.nodeList))
			if len(appData.nodeList) > appData.currentListIndex {
				len := len(appData.nodeList[appData.currentListIndex].Data)
				if len < 16 {
					return 1, len
				} else {
					return len / 16, 16
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
			//fmt.Printf("Current %d Row: %d Col %d\n", appData.currentListIndex, i.Row, i.Col)
			if len(appData.nodeList) > appData.currentListIndex {
				idx := i.Row*16 + i.Col
				if idx < len(appData.nodeList[appData.currentListIndex].Data) {
					o.(*widget.Label).SetText(fmt.Sprintf("%02X",
						appData.nodeList[appData.currentListIndex].Data[idx]))
				}
				return
			}
			o.(*widget.Label).SetText("--")
		},
	)

	return table
}
