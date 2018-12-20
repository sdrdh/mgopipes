package mgopipes

import (
	"log"
	"testing"
)

// func Test_getFormattedPipelineString(t *testing.T) {
// 	type args struct {
// 		jsonString string
// 		args       []interface{}
// 	}
// 	type test struct {
// 		name string
// 		args args
// 		want string
// 	}
// 	tests := []test{
// 		test{
// 			name: "slice of string",
// 			args: args{
// 				"",
// 				[]interface{}{"hi", []string{"siddu", "darpan"}},
// 			},
// 			want: "",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := getFormattedPipelineString(tt.args.jsonString, tt.args.args...); got != tt.want {
// 				t.Errorf("getFormattedPipelineString() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

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

func TestGetPipeline(t *testing.T) {
	type args struct {
		jsonString string
		args       []interface{}
	}
	type test struct {
		name    string
		args    args
		wantErr bool
	}
	tests := []test{
		test{
			"test",
			args{
				jsonString: `[{
					"$project": {
					"ch": {
					"$cond": {
					"if": {
					"$eq": ["$channel", "%s"]
					},
					"then": "yes",
					"else": "no"
					}
					}
					}
					}]`,
				args: []interface{}{"CUSTOM"},
			},
			false,
		}, test{
			"test list",
			args{
				jsonString: `[{
					"$project": {
					"ch": {
					"$cond": {
					"if": {
					"$eq": ["$channel", "%l"]
					},
					"then": "yes",
					"else": "no"
					}
					}
					}
					}]`,
				args: []interface{}{[]string{"CUSTOM"}},
			},
			false,
		}, test{
			"test int list",
			args{
				jsonString: `[{
					"$project": {
					"ch": {
					"$cond": {
					"if": {
					"$eq": ["$channel", "%l"]
					},
					"then": "yes",
					"else": "no"
					}
					}
					}
					}]`,
				args: []interface{}{[]int{1, 2, 3}},
			},
			false,
		}, test{
			"test int list",
			args{
				jsonString: `[{
					"$project": {
					"ch": {
					"$cond": {
					"if": {
					"$eq": ["$channel", "%l"]
					},
					"then": "yes",
					"else": "no"
					}
					}
					}
					}]`,
				args: []interface{}{[]float32{1.0, 2, 3}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPipeline(tt.args.jsonString, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPipeline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println(got)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetPipeline() = %v, want %v", got, tt.want)
			// }
		})
	}
}
