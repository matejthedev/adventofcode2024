package main

import (
	_ "embed"
	"testing"
)

//go:embed data.txt
var dataTxt string

func Test_countAntiNodes(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "countAntiNodes with example data",
			args: args{
				input: "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............",
			},
			want: 14,
		},
		{
			name: "countAntiNodes with task data",
			args: args{
				input: dataTxt,
			},
			want: 320,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAntiNodes(tt.args.input); got != tt.want {
				t.Errorf("countAntiNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countUniqueLocations(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "countUniqueLocations with example data",
			args: args{
				input: "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............",
			},
			want: 34,
		},
		{
			name: "countUniqueLocations with task data",
			args: args{
				input: dataTxt,
			},
			want: 1157,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueLocations(tt.args.input); got != tt.want {
				t.Errorf("countUniqueLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}
