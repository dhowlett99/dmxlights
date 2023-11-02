package buttons

// getNextMenuItem get the next items in the menu sequence.
// Wraps if your at the end.
func getNextMenuItem(currentMode int, chaser bool, editStaticColorMode bool) int {

	menuOrder := []int{NORMAL, NORMAL_STATIC, FUNCTION, CHASER_DISPLAY, CHASER_DISPLAY_STATIC, CHASER_FUNCTION, STATUS}

	if !chaser && !editStaticColorMode {
		switch {
		case currentMode == NORMAL:
			return menuOrder[FUNCTION]

		case currentMode == FUNCTION:
			return menuOrder[STATUS]

		case currentMode == STATUS:
			return menuOrder[NORMAL]
		}
	}

	if !chaser && editStaticColorMode {
		switch {
		case currentMode == NORMAL:
			return menuOrder[NORMAL_STATIC]

		case currentMode == NORMAL_STATIC:
			return menuOrder[FUNCTION]

		case currentMode == FUNCTION:
			return menuOrder[STATUS]

		case currentMode == STATUS:
			return menuOrder[NORMAL]
		}
	}

	if chaser && !editStaticColorMode {
		switch {
		case currentMode == NORMAL:
			return menuOrder[FUNCTION]

		case currentMode == FUNCTION:
			return menuOrder[CHASER_DISPLAY]

		case currentMode == CHASER_DISPLAY:
			return menuOrder[CHASER_FUNCTION]

		case currentMode == CHASER_FUNCTION:
			return menuOrder[STATUS]

		case currentMode == STATUS:
			return menuOrder[NORMAL]
		}
	}

	if chaser && editStaticColorMode {
		switch {
		case currentMode == NORMAL:
			return menuOrder[FUNCTION]

		case currentMode == FUNCTION:
			return menuOrder[CHASER_DISPLAY]

		case currentMode == CHASER_DISPLAY:
			return menuOrder[CHASER_DISPLAY_STATIC]

		case currentMode == CHASER_DISPLAY_STATIC:
			return menuOrder[CHASER_FUNCTION]

		case currentMode == CHASER_FUNCTION:
			return menuOrder[STATUS]

		case currentMode == STATUS:
			return menuOrder[NORMAL]
		}
	}

	return 0
}
