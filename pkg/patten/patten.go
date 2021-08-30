package patten

import "github.com/dhowlett99/dmxlights/pkg/common"

const (
	full = 255
)

func MakePatterns() map[string]common.Patten {

	Pattens := make(map[string]common.Patten)

	standard := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 85, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 170, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
				},
			},
		},
	}

	rgbchase := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}, {R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}, {R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}, {R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}, {R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}, {R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}, {R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}, {R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}, {R: 0, G: 0, B: 0}}},
				},
			},
		},
	}

	pairs := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
		},
	}

	colors := common.Patten{
		Steps: []common.Step{
			{ // Step 1, - Red
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 2 - Orange
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 155, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 3 - Yellow
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 4 - Green
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 5 - Pastel Green
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 50}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 6 - Cyan Blue
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 155, B: 255}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 7 - Blue
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Step 8 - Purple
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}},
				},
			},
			// { // Step 9, - All off
			// 	Fixtures: []common.Fixture{
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 	},
			// },
		},
	}

	singlestep := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			// {
			// 	Fixtures: []common.Fixture{
			// 		{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		// {MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 	},
			// },
		},
	}

	fade := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
		},
	}

	scanner := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
				},
			},
		},
	}

	inverted := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 255},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}, Gobo: 36, Shutter: 255, Pan: 50, Tilt: 200},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 100, Tilt: 150},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 255, B: 255}}, Gobo: 36, Shutter: 255, Pan: 150, Tilt: 100},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
					{MasterDimmer: full, Colors: []common.Color{{R: 100, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 200, Tilt: 50},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 190, B: 255}}, Gobo: 36, Shutter: 255, Pan: 255, Tilt: 0},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
					{MasterDimmer: full, Colors: []common.Color{{R: 145, G: 0, B: 255}}, Gobo: 36, Shutter: 255, Pan: 0, Tilt: 0},
				},
			},
		},
	}

	Pattens["singlestep"] = singlestep
	Pattens["inverted"] = inverted
	Pattens["fade"] = fade
	Pattens["scanner"] = scanner
	Pattens["standard"] = standard
	Pattens["rgbchase"] = rgbchase
	Pattens["pairs"] = pairs
	Pattens["colors"] = colors

	return Pattens
}
