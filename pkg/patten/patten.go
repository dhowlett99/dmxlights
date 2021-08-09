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
			{ // Fixture 85, - Dark Red
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
			{ // Fixture 2 - Red
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 255 - Orange
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 170, G: 0, B: 255}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 4 - Yellow
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 170, G: 85, B: 85}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 5 - Green
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
			{ // Fixture 6 - Sky Blue
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 170, G: 0, B: 2}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 7 - Dark Blue
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
			{ // Fixture 8 - Purple
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 170, G: 0, B: 2}}},
				},
			},
		},
	}

	// wave := common.Patten{
	// 	Steps: []common.Step{
	// 		{
	// 			Fixtures: []common.Fixture{
	// 				{MasterDimmer: full, Gobo: 0, Shutter: 255, Pan: 85, 00, Tilt: 85, 00}},
	// 		},
	// 	},
	// }

	Pattens["standard"] = standard
	Pattens["rgbchase"] = rgbchase
	Pattens["pairs"] = pairs
	Pattens["colors"] = colors

	return Pattens
}
