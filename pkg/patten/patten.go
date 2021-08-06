package patten

import "github.com/dhowlett99/dmxlights/pkg/common"

const (
	full = 3
)

func MakePatterns() map[string]common.Patten {

	Pattens := make(map[string]common.Patten)

	standard := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
				},
			},
		},
	}

	rgbchase := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}, {R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}, {R: 0, G: 3, B: 0}, {R: 0, G: 0, B: 3}}},
				},
			},
		},
	}

	pairs := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
				},
			},
		},
	}

	colors := common.Patten{
		Steps: []common.Step{
			{ // Fixture 1 - Dark Red
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 3, G: 2, B: 1}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 2 - Red
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 3, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 3 - Orange
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 2, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 4 - Yellow
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 2, G: 1, B: 1}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 5 - Green
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 6 - Sky Blue
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 2, G: 0, B: 2}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 7 - Dark Blue
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 3}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{ // Fixture 8 - Purple
				Fixtures: []common.Fixture{
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{Brightness: full, Colors: []common.Color{{R: 2, G: 0, B: 2}}},
				},
			},
		},
	}

	Pattens["standard"] = standard
	Pattens["rgbchase"] = rgbchase
	Pattens["pairs"] = pairs
	Pattens["colors"] = colors

	return Pattens
}
