package editor

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewChannelPanel(channelList []itemSelect, channelOptions []string) *widget.List {

	// Channel or Switch State Selection Panel.
	var channelPanel *widget.List
	channelPanel = widget.NewList(
		// Function to find length.
		func() int {
			return len(channelList)
		},
		// Function to create item.
		func() (o fyne.CanvasObject) {
			return container.NewHBox(
				widget.NewLabel("template"),

				widget.NewSelect(channelOptions, func(value string) {
					lastChannel, _ := strconv.Atoi(o.(*fyne.Container).Objects[0].(*widget.Label).Text)
					//fmt.Printf("We just pressed channel %d and set it to %s\n", lastChannel, value)
					item := itemSelect{}
					item.Number = int16(lastChannel)
					item.Label = value
					item.Options = channelOptions
					channelList = UpdateItem(channelList, item.Number, item)
				}),

				widget.NewButton("-", func() {
					//log.Println("Delete Button pressed for ")
				}),
				widget.NewButton("+", func() {
					//log.Println("Add Button pressed for ")
				}),
			)
		},
		// Function to update item in this list.
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", channelList[i].Number))

			// find the selected option in the options list.
			for _, option := range channelList[i].Options {
				if option == channelList[i].Label {
					o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(option)
				}
			}

			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				channelList = DeleteItem(channelList, channelList[i].Number)
				channelPanel.Refresh()
			}

			o.(*fyne.Container).Objects[3].(*widget.Button).OnTapped = func() {
				channelList = AddItem(channelList, channelList[i].Number, channelOptions)
				channelPanel.Refresh()
			}

		})

	return channelPanel
}
