// Copyright (C) 2022, 2023, 2024, 2025 dhowlett99.
// This is button processor, used by the launchpad and gui interfaces.
// This file processes the error messages.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package buttons

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func displayErrorPopUp(w fyne.Window, errorMessage string) (modal *widget.PopUp) {

	title := widget.NewLabel("Error")
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	errorText := widget.NewLabel(errorMessage)

	// Ok button.
	button := widget.NewButton("Dismiss", func() {
		modal.Hide()
	})

	// Layout of settings panel.
	modal = widget.NewModalPopUp(
		container.NewVBox(
			title,
			errorText,
			widget.NewLabel(""),
			container.NewHBox(layout.NewSpacer(), button),
		),
		w.Canvas(),
	)

	modal.Resize(fyne.NewSize(250, 250))
	modal.Show()

	return modal
}
