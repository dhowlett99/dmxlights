package patten

import (
	"fmt"
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

// GeneratePatten takes an array of coordinates and turns them into a patten
// which is the starting point for all sequence steps.
func GeneratePatten(coordinates []coordinate, NumberFixtures int, shift int) common.Patten {

	// Shift this many for each fixture.
	shiftThisMany := make(map[int]int)

	// The current shift for this fixture
	thisShift := make(map[int]int)

	// Initalise a shift value for each fixture.
	for fx := 0; fx < 4; fx++ {
		thisShift[fx] = 0
	}

	fmt.Printf("NumberFixtures:%d shiftThisMany:%d shift:%d\n", NumberFixtures, shiftThisMany, shift)

	fmt.Printf("=========== S T A R T ===============\n")

	// First create the patten.
	patten := common.Patten{}

	steps := []common.Step{}

	// Now create the steps in the patten.
	for coordinateNumber := range coordinates {

		//fmt.Printf("Coordinate Number:%d\n", coordinateNumber)
		fixtures := []common.Fixture{}

		for fixture := 0; fixture < NumberFixtures; fixture++ {

			// Share the shiftable position out across the number of available fixtures
			shiftThisMany[fixture] = ((len(coordinates)) / 4)

			// Only shift above 0
			if shift > 0 && fixture != 0 {
				// The shiftThisMany is basically a quarter of the availale
				// coordinates, if we multiple by the shift we wiil
				// get either a 1/4 rotation for a shift of 1, 1/2 shift for 2 and 3/4 for 3.
				//fmt.Printf("fixture=%d shiftThisMany=%d  shift=%d Shifting by=%d\n", fixture, shiftThisMany, shift, (shiftThisMany[fixture]*(fixture+1))*shift)
				thisShift[fixture] = (shiftThisMany[fixture] * fixture) * shift
			}
			//fmt.Printf("Coordinate No %d\n", coordinateNumber)
			current := coordinateNumber + thisShift[fixture]

			if current > len(coordinates) {
				current = 1
			}
			if current == len(coordinates) {
				current = 0
			}
			// fmt.Printf("Current No %d\n", current)
			// fmt.Printf("Pan:%d  \n", coordinates[current].Pan)
			// fmt.Printf("Tilt:%d \n", coordinates[current].Tilt)
			// fmt.Printf("thisShift%d\n", thisShift)

			newFixture := common.Fixture{
				Type:         "scanner",
				MasterDimmer: full,
				Colors: []common.Color{
					common.GetColorButtonsArray(current),
				},
				Pan:     int(coordinates[current].Pan),
				Tilt:    int(coordinates[current].Tilt),
				Shutter: 255,
				Gobo:    36,
			}
			fixtures = append(fixtures, newFixture)

			//fmt.Printf("Pan:%d Tilt:%d\n", int(coordinates[current].Pan), int(coordinates[current].Tilt))

			//fmt.Printf("==================================\n")
		}
		newStep := common.Step{
			Type:     "scanner",
			Fixtures: fixtures,
		}
		steps = append(steps, newStep)
	}
	patten.Name = "circle"
	patten.Steps = steps

	for stepNumber, step := range steps {

		fmt.Printf("Step %d\n", stepNumber)
		for _, fixture := range step.Fixtures {
			fmt.Printf("Fixture Pan %d  Tilt %d \n", fixture.Pan, fixture.Tilt)
		}
	}

	return patten

}

type coordinate struct {
	Tilt int
	Pan  int
}

func CircleGenerator(size int) (out []coordinate) {
	var theta float64
	for theta = 0; theta <= 360; theta += 10 {
		n := coordinate{}
		n.Tilt, n.Pan = circleXY(float64(size), theta)
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
		n.Tilt = int(x)
		n.Pan = int(t)
		out = append(out, n)
	}
	return out
}

func ScanGeneratorUpDown(size int) (out []coordinate) {
	pan := 128
	for tilt := 0; tilt < 255; tilt += 10 {
		n := coordinate{}
		n.Tilt = tilt
		n.Pan = pan
		out = append(out, n)
	}
	return out
}

func ScanGeneratorLeftRight(size int) (out []coordinate) {
	tilt := 128
	for pan := 0; pan < 255; pan += 10 {
		n := coordinate{}
		n.Tilt = tilt
		n.Pan = pan
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
