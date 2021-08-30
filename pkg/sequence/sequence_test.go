package sequence

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_translatePatten(t *testing.T) {
	type args struct {
		steps []common.Step
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		// {
		// 	name: "simple one step, fade up one lamp and then fade it down.",
		// 	args: args{
		// 		steps: []common.Step{
		// 			{ // This input step cause the translate to add 7 fade up steps.
		// 				Fixtures: []common.Fixture{
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
		// 				},
		// 			},
		// 			{ // This input step cause the translate to add 7 fade down steps.
		// 				Fixtures: []common.Fixture{
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: []common.Step{
		// 		{ // Step 0
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 1
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 2
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 3
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 4
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 5
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 6 full brightness.
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 7 full brightness. Start of fade down.
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 8
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 9
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{ // Step 10
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 189, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 	},
		// },
		{

			name: "eight lamps and then simple 1 - 2 fade up and down.",
			args: args{
				steps: []common.Step{
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
					{
						Fixtures: []common.Fixture{
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
							{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						},
					},
				},
			},
			want: []common.Step{
				{
					Fixtures: []common.Fixture{
						{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
		},
		// {
		// 	name: "now operate on 8 lamps.",
		// 	args: args{
		// 		steps: []common.Step{
		// 			{
		// 				Fixtures: []common.Fixture{
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				},
		// 			},
		// 			{
		// 				Fixtures: []common.Fixture{
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 					{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: []common.Step{
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 66, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 127, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 180, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 220, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 246, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 		{
		// 			Fixtures: []common.Fixture{
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 255, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 				{MasterDimmer: 255, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
		// 			},
		// 		},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translatePatten(tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				printSteps(got)
				t.Errorf("got= %+v", got)
				//printSteps(tt.want)
				t.Errorf("want %+v", tt.want)
			}
		})
	}
}

func printSteps(steps []common.Step) {

	fmt.Println()
	for stepIndex, step := range steps {
		fmt.Printf("Step No:%d\n", stepIndex)
		for fixtureIndex, fixture := range step.Fixtures {
			fmt.Printf("\t\tFixture No:%d\n", fixtureIndex)
			for _, color := range fixture.Colors {
				fmt.Printf("\t\t\tColor   R:%d G:%d B:%d\n", color.R, color.G, color.B)
			}
		}
	}
}
