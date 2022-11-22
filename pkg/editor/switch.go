package editor

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SwitchPanel struct {
	SwitchPanel   *widget.List
	SwitchesList  []itemSelect
	SwitchOptions []string
}

func NewSwitchPanel(switchesAvailable bool, switchesList []itemSelect, switchOptions []string, ap *ActionPanel) *SwitchPanel {

	sw := SwitchPanel{}
	sw.SwitchesList = switchesList
	sw.SwitchOptions = switchOptions

	// Switches Selection Panel.
	if switchesAvailable {
		sw.SwitchPanel = widget.NewList(
			func() int {
				return len(sw.SwitchesList)
			},
			func() (o fyne.CanvasObject) {
				return container.NewHBox(
					widget.NewLabel("template"),

					widget.NewSelect(sw.SwitchOptions, func(value string) {
						//label := o.(*fyne.Container).Objects[0].(*widget.Label).Text
						//fmt.Printf("Select set to %s number %s\n", value, label)
					}),

					widget.NewButton("Select", nil),
				)
			},
			// Function to update item in this list.
			func(i widget.ListItemID, o fyne.CanvasObject) {
				//fmt.Printf("Switch ID is %d   Switch Setting is %s\n", i, switchOptions[i])
				o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", sw.SwitchesList[i].Number))

				// find the selected option in the options list.
				for _, option := range sw.SwitchesList[i].Options {
					if option == sw.SwitchesList[i].Label {
						o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(option)
					}
				}

				// new part
				o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
					ap.ActionsList = []actionItems{}
					for _, action := range sw.SwitchesList[i].Actions {
						newAction := actionItems{}
						newAction.Name = action.Name
						newAction.Colors = action.Colors
						newAction.Mode = action.Mode
						newAction.Fade = action.Fade
						newAction.Speed = action.Speed
						ap.ActionsList = append(ap.ActionsList, newAction)
					}
					fmt.Printf("I am button %d actions %+v\n", sw.SwitchesList[i].Number, ap.ActionsList)
					ap.ActionsPanel.Hidden = false
					ap.ActionsPanel.Refresh()
				}
			})
	}
	return &sw
}
