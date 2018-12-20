package mgopipes

import (
	"testing"
)

func Test_getFormattedPipelineString(t *testing.T) {
	type args struct {
		jsonString string
		args       []interface{}
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		test{
			name: "slice of string",
			args: args{
				"",
				[]interface{}{"hi", [2]string{"siddu", "darpan"}},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFormattedPipelineString(tt.args.jsonString, tt.args.args...); got != tt.want {
				t.Errorf("getFormattedPipelineString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceOrArray(t *testing.T) {
	type test struct {
		name string
		arg  interface{}
		want bool
	}
	tests := []test{
		test{
			"Slice",
			[]int{1, 2, 3, 4, 5},
			true,
		},
		test{
			"array",
			[2]int{1, 2},
			true,
		},
		test{
			"int",
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceOrArray(tt.arg); got != tt.want {
				t.Errorf("sliceOrArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
