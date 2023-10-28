package buttons

// Create a sequence of menu items based on if chaser is running.
func newMenu(chaser bool) []int {

	if !chaser {
		return []int{NORMAL, FUNCTION, STATUS}
	} else {
		return []int{NORMAL, FUNCTION, CHASER_DISPLAY, CHASER_FUNCTION, STATUS}
	}

}

// getNextMenuItem get the next items in the menu sequence.
// Wraps if your at the end.
func getNextMenuItem(selectedMode int, chaser bool) int {

	menuOrder := newMenu(chaser)

	if !chaser {
		switch {
		case selectedMode == NORMAL:
			return menuOrder[1]

		case selectedMode == FUNCTION:
			return menuOrder[2]

		case selectedMode == STATUS:
			return menuOrder[0]
		}
	}

	if chaser {
		switch {
		case selectedMode == NORMAL:
			return menuOrder[1]

		case selectedMode == FUNCTION:
			return menuOrder[2]

		case selectedMode == CHASER_DISPLAY:
			return menuOrder[3]

		case selectedMode == CHASER_FUNCTION:
			return menuOrder[4]

		case selectedMode == STATUS:
			return menuOrder[0]
		}
	}

	return 0
}
