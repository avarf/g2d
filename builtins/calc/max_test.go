package calc

import (
	"fmt"
	"testing"

	"github.com/lucasepe/g2d/object"
)

func TestMaxOfArray(t *testing.T) {
	cases := []struct {
		input []object.Object
		want  object.Object
	}{
		{
			input: []object.Object{
				&object.Float{Value: 2},
				&object.Float{Value: -22.45},
				&object.Integer{Value: 10},
				&object.Float{Value: 2},
				&object.Integer{Value: -89},
				&object.Float{Value: 6.876},
			},
			want: &object.Float{Value: 10},
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("max_%d", i), func(t *testing.T) {
			got := Max(nil, tt.input...)
			if got.(object.Comparable).Compare(tt.want) != 0 {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}
