package editor

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ChannelPanel struct {
	ChannelPanel   *widget.List
	ChannelList    []itemSelect
	ChannelOptions []string
}

func NewChannelPanel(channelList []itemSelect, channelOptions []string) *ChannelPanel {

	cp := ChannelPanel{}
	cp.ChannelList = channelList
	cp.ChannelOptions = channelOptions

	// Channel or Switch State Selection Panel.
	cp.ChannelPanel = widget.NewList(
		// Function to find length.
		func() int {
			return len(cp.ChannelList)
		},
		// Function to create item.
		func() (o fyne.CanvasObject) {
			return container.NewHBox(
				widget.NewLabel("template"),

				widget.NewSelect(cp.ChannelOptions, func(value string) {
					lastChannel, _ := strconv.Atoi(o.(*fyne.Container).Objects[0].(*widget.Label).Text)
					//fmt.Printf("We just pressed channel %d and set it to %s\n", lastChannel, value)
					item := itemSelect{}
					item.Number = int16(lastChannel)
					item.Label = value
					item.Options = cp.ChannelOptions
					cp.ChannelList = UpdateItem(cp.ChannelList, item.Number, item)
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
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", cp.ChannelList[i].Number))

			// find the selected option in the options list.
			for _, option := range cp.ChannelList[i].Options {
				if option == cp.ChannelList[i].Label {
					o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(option)
				}
			}

			o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
				cp.ChannelList = DeleteItem(cp.ChannelList, cp.ChannelList[i].Number)
				cp.ChannelPanel.Refresh()
			}

			o.(*fyne.Container).Objects[3].(*widget.Button).OnTapped = func() {
				cp.ChannelList = AddItem(cp.ChannelList, cp.ChannelList[i].Number, cp.ChannelOptions)
				cp.ChannelPanel.Refresh()
			}

		})

	return &cp
}
