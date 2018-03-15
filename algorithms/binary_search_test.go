package algorithms

import "testing"

func genIntAry(cap int) (res []int) {
	for i := 1; i <= cap; i++ {
		res = append(res, i)
	}
	return
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		list []int
		item int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "",
			args:    args{[]int{2, 4, 6, 8, 10}, 6},
			want:    2,
			wantErr: false,
		},
		{
			name:    "",
			args:    args{[]int{2, 4, 6, 8, 10}, 7},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "",
			args:    args{genIntAry(1000), 789},
			want:    788,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearch(tt.args.list, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinarySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
