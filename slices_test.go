package moreslices

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		s     []int
		e     int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Insert tail",
			args: args{
				[]int{1, 2, 3, 4},
				5,
				4,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Insert head",
			args: args{
				[]int{1, 2, 3, 4},
				5,
				0,
			},
			want: []int{5, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.s, tt.args.e, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteAll(t *testing.T) {
	type args struct {
		s     []int
		pFunc func(e int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Delete2",
			args: args{
				[]int{1, 2, 2, 3},
				func(e int) bool {
					return e == 2
				},
			},
			want: []int{1, 3},
		},
		{
			name: "DeleteEmpty",
			args: args{
				[]int{1, 3},
				func(e int) bool {
					return e == 2
				},
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteAll(tt.args.s, tt.args.pFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
