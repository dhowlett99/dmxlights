package editor

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ActionPanel struct {
	ActionsPanel  *widget.List
	ActionsList   []itemSelect
	ActionOptions []string
}

func NewActionsPanel(actionsAvailable bool, actionsList []itemSelect, actionOptions []string) *ActionPanel {

	ap := ActionPanel{}
	ap.ActionsList = actionsList
	ap.ActionOptions = actionOptions

	// Actions Selection Panel.
	var actionsPanel *widget.List
	if actionsAvailable {
		ap.ActionsPanel = widget.NewList(
			func() int {
				return len(ap.ActionsList)
			},
			func() fyne.CanvasObject {
				return container.NewHBox(
					widget.NewLabel("template"),

					widget.NewSelect(ap.ActionOptions, func(value string) {
						//fmt.Printf("Select action set to %s\n", value)
					}),

					widget.NewButton("-", func() {
						//fmt.Printf("Delete action Button pressed for action\n")
					}),
					widget.NewButton("+", func() {
						//fmt.Printf("Add action Button pressed for action\n")
					}),
				)
			},
			// Function to update item in this list.
			func(i widget.ListItemID, o fyne.CanvasObject) {
				//fmt.Printf("Action ID is %d   Action Setting is %s\n", i, actionOptions[i])
				o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%d", ap.ActionsList[i].Number))

				// find the selected option in the options list.
				//for _, option := range actionsList[i].Options {
				//if option == actionsList[i].Label {
				fmt.Printf("---> Found %s\n", ap.ActionsList[i].Label)
				o.(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Label)
				//}
				//}

				o.(*fyne.Container).Objects[2].(*widget.Button).OnTapped = func() {
					ap.ActionsList = DeleteItem(ap.ActionsList, ap.ActionsList[i].Number)
					actionsPanel.Refresh()
				}

				o.(*fyne.Container).Objects[3].(*widget.Button).OnTapped = func() {
					ap.ActionsList = AddItem(ap.ActionsList, ap.ActionsList[i].Number, ap.ActionOptions)
					actionsPanel.Refresh()
				}
			})
	}
	return &ap
}

func (ap *ActionPanel) UpdateActionList(actionList []itemSelect) {

	ap.ActionsList = actionList

}
