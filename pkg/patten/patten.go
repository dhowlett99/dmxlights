package patten

import (
	"math"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

const (
	full = 255
)

func MakePatterns() map[string]common.Patten {

	Pattens := make(map[string]common.Patten)

	standard := common.Patten{
		Steps: []common.Step{
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
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}, {R: 0, G: 255, B: 0}, {R: 0, G: 0, B: 255}}},
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
			// {
			// 	Fixtures: []common.Fixture{
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
			// 		{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
			// 	},
			// },
		},
	}

	inward := common.Patten{
		Steps: []common.Step{
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
				},
			},
			{
				Fixtures: []common.Fixture{
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					{MasterDimmer: full, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
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

	Pattens["standard"] = standard
	Pattens["rgbchase"] = rgbchase
	Pattens["pairs"] = pairs
	Pattens["inward"] = inward
	Pattens["colors"] = colors

	return Pattens

}

func GeneratePatten(coordinates []coordinate, NumberFixtures int, shift bool) common.Patten {

	reverseCoordinates := []coordinate{}

	for c := len(coordinates); c > 0; c-- {
		coordinate := coordinates[c-1]
		reverseCoordinates = append(reverseCoordinates, coordinate)
	}

	// First create the patten.
	patten := common.Patten{}

	steps := []common.Step{}

	// Now create the steps in the patten.
	for index, coordinate := range coordinates {
		fixtures := []common.Fixture{}
		for f := 0; f < NumberFixtures; f++ {

			if f == 0 || !shift {
				newFixture := common.Fixture{
					Type:         "scanner",
					MasterDimmer: full,
					Colors: []common.Color{
						{
							R: int(coordinate.y),
							G: int(coordinate.x),
							B: makeBlue(int(coordinate.y), int(coordinate.x)),
						},
					},
					Pan:     int(coordinate.y),
					Tilt:    int(coordinate.x),
					Shutter: 255,
					Gobo:    36,
				}
				fixtures = append(fixtures, newFixture)
			}
			if f == 1 && shift {
				newFixture := common.Fixture{
					Type:         "scanner",
					MasterDimmer: full,
					Colors: []common.Color{
						{
							R: int(reverseCoordinates[index].y),
							G: int(reverseCoordinates[index].x),
							B: 0,
						},
					},
					Pan:     int(reverseCoordinates[index].y),
					Tilt:    int(reverseCoordinates[index].x),
					Shutter: 255,
					Gobo:    36,
				}
				fixtures = append(fixtures, newFixture)
			}

		}
		newStep := common.Step{
			Type:     "scanner",
			Fixtures: fixtures,
		}
		steps = append(steps, newStep)
	}
	patten.Name = "circle"
	patten.Steps = steps

	return patten

}

func makeBlue(red, green int) int {

	return red + green/2
}

type coordinate struct {
	x int
	y int
}

func CircleGenerator(size int) (out []coordinate) {
	var theta float64
	for theta = 0; theta <= 360; theta += 10 {
		n := coordinate{}
		n.x, n.y = circleXY(float64(size), theta)
		out = append(out, n)
	}
	return out
}

func ScanGenerateSineWave(size int, frequency int) (out []coordinate) {
	var t float64
	T := float64(size)
	for t = 1; t < T-1; t += 10 {
		n := coordinate{}
		x := (float64(size)/2 + math.Sin(t*float64(frequency))*100)
		n.x = int(x)
		n.y = int(t)
		out = append(out, n)
	}
	return out
}

func ScanGeneratorUpDown(size int) (out []coordinate) {
	pan := 128
	for tilt := 0; tilt < 255; tilt += 10 {
		n := coordinate{}
		n.x = tilt
		n.y = pan
		out = append(out, n)
	}
	return out
}

func ScanGeneratorLeftRight(size int) (out []coordinate) {
	tilt := 128
	for pan := 0; pan < 255; pan += 10 {
		n := coordinate{}
		n.x = tilt
		n.y = pan
		out = append(out, n)
	}
	return out
}

func circleXY(r float64, theta float64) (int, int) {
	// Convert angle to radians
	theta = (theta - 90) * math.Pi / 180

	x := int(r*math.Cos(theta) + 128)
	y := int(-r*math.Sin(theta) + 128)
	return x, y
}
