package solutions

import "testing"

func Test_canPartitions(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test", args: args{nums: []int{1, 5, 11, 5}}, want: true},
		{name: "test", args: args{nums: []int{1, 2, 3, 5}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
