package main

import "testing"

var reports = parseData()

func Test_CountSafeReports(t *testing.T) {
	type args struct {
		reports [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example test",
			args: args{
				reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			want: 2,
		},
		{
			name: "real data test",
			args: args{
				reports,
			},
			want: 218,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSafeReports(tt.args.reports); got != tt.want {
				t.Errorf("countSafeReports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSafeReportsWithTolerate(t *testing.T) {
	type args struct {
		reports [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example test with tolerate",
			args: args{
				reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			want: 4,
		},
		{
			name: "real data test with tolerate",
			args: args{
				reports,
			},
			want: 290,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSafeReportsWithTolerate(tt.args.reports); got != tt.want {
				t.Errorf("countSafeReportsWithTolerate() = %v, want %v", got, tt.want)
			}
		})
	}
}
