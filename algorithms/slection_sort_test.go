package algorithms

import (
	"reflect"
	"testing"
)

func Test_findSmallest(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name         string
		args         args
		wantSmallest [2]int
	}{
		{
			name:         "",
			args:         args{[]int{5, 32, 5, 1, 68, 2, 44, 2}},
			wantSmallest: [2]int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSmallest := findSmallest(tt.args.in); !reflect.DeepEqual(gotSmallest, tt.wantSmallest) {
				t.Errorf("findSmallest() = %v, want %v", gotSmallest, tt.wantSmallest)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []int
	}{
		{
			name:    "",
			args:    args{[]int{5, 32, 5, 1, 68, 2, 44, 2}},
			wantOut: []int{1, 2, 2, 5, 5, 32, 44, 68},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := SelectionSort(tt.args.in); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("SelectionSort() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
