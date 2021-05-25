package sequence

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dhowlett99/dmxlights/pkg/common"
)

func Test_calcSteps(t *testing.T) {

	type args struct {
		lastStep common.Step
		nextStep common.Step
	}
	tests := []struct {
		name string
		args args
		want []common.Step
	}{
		{
			name: "pass",
			args: args{
				lastStep: common.Step{
					Fixtures: []common.Fixture{
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				nextStep: common.Step{
					Fixtures: []common.Fixture{
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
			want: []common.Step{
				{
					Fixtures: []common.Fixture{
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 3, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 2, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 1, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
				{
					Fixtures: []common.Fixture{
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
						{Brightness: 3, Colors: []common.Color{{R: 0, G: 0, B: 0}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcSteps(tt.args.lastStep, tt.args.nextStep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcSteps() want %v", tt.want)
				t.Errorf("calcSteps() got  %v", got)
				for _, step := range tt.want {
					for fixure := 0; fixure < len(step.Fixtures); fixure++ {
						fmt.Printf("Color  R:%d G:%d B:%d\n",
							step.Fixtures[fixure].Colors[0].R,
							step.Fixtures[fixure].Colors[0].G,
							step.Fixtures[fixure].Colors[0].R)
					}
				}
				fmt.Println()
				fmt.Println("----------------------------")
				fmt.Println()
				for _, step := range got {
					for fixure := 0; fixure < len(step.Fixtures); fixure++ {
						fmt.Printf("Color  R:%d G:%d B:%d\n",
							step.Fixtures[fixure].Colors[0].R,
							step.Fixtures[fixure].Colors[0].G,
							step.Fixtures[fixure].Colors[0].R)
					}
				}
			}
		})
	}
}
