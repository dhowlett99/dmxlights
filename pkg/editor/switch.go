package editor

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewSwitchPanel(switchesAvailable bool, switchesList []itemSelect, switchOptions []string, ap *ActionPanel) *widget.List {
	// Switches Selection Panel.
	var switchesPanel *widget.List
	if switchesAvailable {
		switchesPanel = widget.NewList(
			func() int {
				return len(switchesList)
			},
			func() (o fyne.CanvasObject) {
				return container.NewHBox(
					widget.NewLabel("template"),

					widget.NewSelect(switchOptions, func(value string) {
						//label := o.(*fyne.Container).Objects[0].(*widget.Label).Text
						//fmt.Printf("Select set to %s number %s\n", value, label)
					}),

					widget.NewButton("Select", nil),
				)
			},
			// Function to update item in this list.
			func(i widget.ListItemID, o fyne.CanvasObject) {
				//fmt.Printf("Switch ID is %d   Switch Setting is %s\n", i, switchOptions[i])
				o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", switchesList[i].Number))

				// find the selected option in the options list.
				for _, option := range switchesList[i].Options {
					if option == switchesList[i].Label {
						o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(option)
					}
				}

				// new part
				o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
					ap.ActionsList = []itemSelect{}
					for _, action := range switchesList[i].Actions {
						newAction := itemSelect{}
						newAction.Label = action.Label
						ap.ActionsList = append(ap.ActionsList, newAction)
					}
					fmt.Printf("I am button %d actions %+v\n", switchesList[i].Number, ap.ActionsList)
					ap.ActionsPanel.Refresh()
				}
			})
	}
	return switchesPanel
}
