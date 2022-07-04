package flash

import (
	"time"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

type Flash struct {
	Enabled    bool
	X          int
	Y          int
	onColor    common.Color
	offColor   common.Color
	guiButtons chan common.ALight
}

func NewFlashButton(X int, Y int, onColor common.Color, offColor common.Color, guiButtons chan common.ALight) Flash {

	f := Flash{
		Enabled:    false,
		X:          X,
		Y:          Y,
		onColor:    onColor,
		offColor:   offColor,
		guiButtons: guiButtons,
	}

	return f
}

func StartFlashButton(f *Flash) {

	f.Enabled = true
	go func() {

		for {
			if f.Enabled {
				// Send message to GUI buttons.
				event := common.ALight{
					X:          f.X,
					Y:          f.Y + 1,
					Brightness: 255,
					Red:        f.onColor.R,
					Green:      f.onColor.G,
					Blue:       f.onColor.B,
				}
				f.guiButtons <- event

				time.Sleep(500 * time.Millisecond)

				// Send message to GUI buttons.
				event = common.ALight{
					X:          f.X,
					Y:          f.Y + 1,
					Brightness: 255,
					Red:        f.offColor.R,
					Green:      f.offColor.G,
					Blue:       f.offColor.B,
				}
				f.guiButtons <- event

				time.Sleep(500 * time.Millisecond)
			} else {
				return
			}
		}
	}()
}

func (f *Flash) StopFlashButton() {
	f.Enabled = false
}
