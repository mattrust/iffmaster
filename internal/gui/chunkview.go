// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/encoding/charmap"
)

func NewHexTableView(appData *AppData) *widget.Table {
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
			o.(*widget.Label).SetText("")
		},
	)

	return table
}

func NewIsoTableView(appData *AppData) *widget.Table {
	table := widget.NewTable(
		// Provide the size of the table
		func() (int, int) {
			if len(appData.nodeList) > appData.currentListIndex {
				len := len(appData.nodeList[appData.currentListIndex].Data)
				if len > 0 {
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
					data := appData.nodeList[appData.currentListIndex].Data[idx]
					o.(*widget.Label).SetText(iso8859ToUtf8Char(data))

					return
				}
			}
			o.(*widget.Label).SetText("")
		},
	)

	return table
}

func iso8859ToUtf8Char(isoChar byte) string {
	// handle special characters
	if isoChar == 0 {
		return "\\0"
	} else if isoChar == 9 {
		return "\\t"
	} else if isoChar == 10 {
		return "\\n"
	} else if isoChar == 13 {
		return "\\r"
	} else {
		decoded, err := charmap.ISO8859_1.NewDecoder().Bytes([]byte{isoChar})
		if err != nil {
			return "\\#"
		}
		return string(decoded)
	}
}
