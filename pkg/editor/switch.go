package editor

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dhowlett99/dmxlights/pkg/fixture"
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

func PopulateSwitches(switchOptions []string, fixture fixture.Fixture) (switchesAvailable bool, actionsAvailable bool,
	actionsList []actionItems, switchesList []itemSelect) {

	// Populate switch state settings and actions.
	if fixture.Type == "switch" {
		//labelSwitch.Text = "Switch States"
		for _, state := range fixture.States {
			switchesAvailable = true
			newSelect := itemSelect{}
			newSelect.Number = state.Number
			newSelect.Label = state.Name
			newSelect.Options = switchOptions
			if state.Actions != nil {
				actionsAvailable = true
				actionsList = []actionItems{}
				for _, action := range state.Actions {
					fmt.Printf("----->Add action %+v\n", action)
					newAction := actionItems{}
					newAction.Name = action.Name
					newAction.Colors = strings.Join(action.Colors[:], ",")
					newAction.Mode = action.Mode
					newAction.Fade = action.Fade
					newAction.Speed = action.Speed
					actionsList = append(actionsList, newAction)
				}
			}
			newSelect.Actions = actionsList
			fmt.Printf("----->Actions %+v\n", newSelect.Actions)
			switchesList = append(switchesList, newSelect)
		}
	}

	return switchesAvailable, actionsAvailable, actionsList, switchesList
}
