package main

import "testing"

var slice1, slice2, _ = parseData()

func TestAddUpDistances(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple example of two list",
			args: args{
				s1: []int{1, 3, 2, 4, 5},
				s2: []int{2, 3, 4, 6, 5},
			},
			want: 5,
		},
		{
			name: "real values from the task",
			args: args{
				s1: slice1,
				s2: slice2,
			},
			want: 2815556,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addUpDistances(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("AddUpDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSimilarityScore(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "real values from the task",
			args: args{
				s1: slice1,
				s2: slice2,
			},
			want: 23927637,
		},
		{
			name: "example values for similarity score",
			args: args{
				s1: []int{3, 4, 2, 1, 3, 3},
				s2: []int{4, 3, 5, 3, 9, 3},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSimilarityScore(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("calculateSimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
