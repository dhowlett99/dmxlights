package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type actionItems struct {
	Name   string
	Colors string
	Mode   string
	Fade   string
	Speed  string
}

type ActionPanel struct {
	ActionsPanel        *widget.List
	ActionsList         []actionItems
	ActionNameOptions   []string
	ActionColorsOptions []string
	ActionModeOptions   []string
	ActionFadeOptions   []string
	ActionSpeedOptions  []string
}

func NewActionsPanel(actionsAvailable bool, actionsList []actionItems,
	actionNameOptions []string,
	actionColorsOptions []string,
	actionModeOptions []string,
	actionFadeOptions []string,
	actionSpeedOptions []string) *ActionPanel {

	ap := ActionPanel{}
	ap.ActionsList = actionsList
	ap.ActionNameOptions = actionNameOptions
	ap.ActionColorsOptions = actionColorsOptions
	ap.ActionModeOptions = actionModeOptions
	ap.ActionFadeOptions = actionFadeOptions
	ap.ActionSpeedOptions = actionSpeedOptions

	// Actions Selection Panel.
	if actionsAvailable {
		ap.ActionsPanel = widget.NewList(
			func() int {
				return len(ap.ActionsList)
			},
			func() fyne.CanvasObject {
				return container.NewVBox(
					container.NewHBox(
						widget.NewLabel("Name"),
						widget.NewSelect(ap.ActionNameOptions, func(value string) {}),
					),
					container.NewHBox(
						widget.NewLabel("Colors"),
						widget.NewSelect(ap.ActionColorsOptions, func(value string) {}),
					),
					container.NewHBox(
						widget.NewLabel("Mode"),
						widget.NewSelect(ap.ActionModeOptions, func(value string) {}),
					),
					container.NewHBox(
						widget.NewLabel("Fade"),
						widget.NewSelect(ap.ActionFadeOptions, func(value string) {}),
					),
					container.NewHBox(
						widget.NewLabel("Speed"),
						widget.NewSelect(ap.ActionSpeedOptions, func(value string) {}),
					),
				)
			},
			// Function to update item in this list.
			func(i widget.ListItemID, o fyne.CanvasObject) {
				o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Name)
				o.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Colors)
				o.(*fyne.Container).Objects[2].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Mode)
				o.(*fyne.Container).Objects[3].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Fade)
				o.(*fyne.Container).Objects[4].(*fyne.Container).Objects[1].(*widget.Select).SetSelected(ap.ActionsList[i].Speed)
			})
	}
	return &ap
}

func (ap *ActionPanel) UpdateActionList(actionList []actionItems) {
	ap.ActionsList = actionList
}
