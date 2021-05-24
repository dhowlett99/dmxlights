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

	Pattens["standard"] = standard
	Pattens["rgbchase"] = rgbchase
	Pattens["pairs"] = pairs

	return Pattens
}
