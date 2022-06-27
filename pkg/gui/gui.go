package gui

import (
	"github.com/dhowlett99/dmxlights/pkg/common"
)

func LightLamp(X, Y, R, G, B, Master int, eventsForGui chan common.ALight) {

	// Now trigger the fixture lamp on the launch pad by sending an event.
	e := common.ALight{
		X:          X,
		Y:          Y,
		Brightness: Master,
		Red:        R,
		Green:      G,
		Blue:       B,
	}
	eventsForGui <- e
}
