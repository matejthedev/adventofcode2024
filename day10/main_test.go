package main

import (
	_ "embed"
	"testing"
)

//go:embed data.txt
var dataTxt string

func Test_getTrailHeadsTotalScore(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getTrailHeadsTotalScore with example data",
			args: args{
				input: "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732",
			},
			want: 36,
		},
		{
			name: "getTrailHeadsTotalScore with task data",
			args: args{
				input: dataTxt,
			},
			want: 717,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTrailHeadsTotalScore(tt.args.input); got != tt.want {
				t.Errorf("getTrailHeadsTotalScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTrailheadsRatingSum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getTrailheadsRatingSum with example data",
			args: args{
				input: "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732",
			},
			want: 81,
		},
		{
			name: "getTrailheadsRatingSum with task data",
			args: args{
				input: dataTxt,
			},
			want: 1686,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTrailheadsRatingSum(tt.args.input); got != tt.want {
				t.Errorf("getTrailheadsRatingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
