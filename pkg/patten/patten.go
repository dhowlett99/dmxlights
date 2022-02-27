package patten

import (
	"fmt"
	"math"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

const debug = false

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
		},
	}

	Pattens["standard"] = standard
	Pattens["rgbchase"] = rgbchase
	Pattens["pairs"] = pairs
	Pattens["inward"] = inward
	Pattens["colors"] = colors

	return Pattens

}

// Storage for scanner values.
type scanner struct {
	values []int
}

// GeneratePatten takes an array of coordinates and turns them into a patten
// which is the starting point for all sequence steps.
func GeneratePatten(coordinates []coordinate, NumberFixtures int, requestedShift int, chase bool) common.Patten {

	NumberCoordinates := len(coordinates)

	if debug {
		fmt.Printf("Number Fixtures %d\n", NumberFixtures)
		fmt.Printf("Number Coordinates %d\n", NumberCoordinates)
	}
	// First create the patten.
	patten := common.Patten{}

	steps := []common.Step{}

	// Storage space for the fixtures
	scanners := []scanner{}

	// First generate the values for all posible fixtures ie 8
	for fixture := 0; fixture < 4; fixture++ {

		// new scanner
		s := scanner{}

		actualShift := (NumberCoordinates / 4) * requestedShift

		shift := fixture * actualShift

		if shift == NumberCoordinates {
			shift = 0
		}

		if shift == NumberCoordinates+NumberCoordinates/2 {
			shift = NumberCoordinates / 2
		}

		if shift == (NumberCoordinates*2)+(NumberCoordinates/4) {
			shift = NumberCoordinates / 4
		}
		for coordinate := shift; coordinate < NumberCoordinates; coordinate++ {
			s.values = append(s.values, coordinate)
		}
		for coordinate := 0; coordinate < shift; coordinate++ {
			s.values = append(s.values, coordinate)
		}

		// append the scanner to the list of scanners.
		scanners = append(scanners, s)
	}

	for fixture := 0; fixture < 4; fixture++ {

		// new scanner
		s := scanner{}

		actualShift := (NumberCoordinates / 4) * requestedShift

		shift := fixture * actualShift

		if shift == NumberCoordinates {
			shift = 0
		}

		if shift == NumberCoordinates+NumberCoordinates/2 {
			shift = NumberCoordinates / 2
		}

		if shift == (NumberCoordinates*2)+(NumberCoordinates/4) {
			shift = NumberCoordinates / 4
		}
		for coordinate := shift; coordinate < NumberCoordinates; coordinate++ {
			s.values = append(s.values, coordinate)
		}
		for coordinate := 0; coordinate < shift; coordinate++ {
			s.values = append(s.values, coordinate)
		}

		// append the scanner to the list of scanners.
		scanners = append(scanners, s)
	}

	if debug {
		for _, scanner := range scanners {
			fmt.Printf("scanner %+v\n", scanner)
		}
	}

	var shutterValue int
	// Now create the steps in the patten.
	for f := 0; f < NumberCoordinates; f++ {

		fixtures := []common.Fixture{}

		for step := 0; step < NumberFixtures; step++ {

			if chase { // Flash the scanners in order.
				shutterValue = calulateShutterValue(f, step, NumberFixtures, NumberCoordinates)
			} else {
				shutterValue = 255 // Otherwise just turn on every scanner
			}

			newFixture := common.Fixture{
				Type:         "scanner",
				MasterDimmer: full,
				Colors: []common.Color{
					common.GetColorButtonsArray(scanners[step].values[f]),
				},
				Pan:     int(coordinates[scanners[step].values[f]].Pan),
				Tilt:    int(coordinates[scanners[step].values[f]].Tilt),
				Shutter: shutterValue,
				Gobo:    36,
			}
			fixtures = append(fixtures, newFixture)
		}

		newStep := common.Step{
			Type:     "scanner",
			Fixtures: fixtures,
		}
		steps = append(steps, newStep)
		patten.Name = "circle"
		patten.Steps = steps
	}
	return patten
}

func calulateShutterValue(currentCoordinate int, currentStep int, NumberFixtures int, NumberCoordinates int) int {

	howOften := NumberCoordinates / NumberFixtures

	if currentCoordinate/howOften == currentStep {
		return 255
	}
	return 0
}

type coordinate struct {
	Tilt int
	Pan  int
}

func CircleGenerator(size int) (out []coordinate) {
	var theta float64
	for theta = 0; theta < 360; theta += 10 {
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
