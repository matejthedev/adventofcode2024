package main

import (
	_ "embed"
	"testing"
)

//go:embed data.txt
var dataTxt string

func Test_getVisitedPositions(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getVisitedPositions with example data",
			args: args{
				input: "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...",
			},
			want: 41,
		},
		{
			name: "getVisitedPositions with task data",
			args: args{
				input: dataTxt,
			},
			want: 5162,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVisitedPositions(tt.args.input); got != tt.want {
				t.Errorf("getVisitedPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumberOfObstructions(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getNumberOfObstructions with example data",
			args: args{
				input: "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...",
			},
			want: 6,
		},
		{
			name: "getNumberOfObstructions with task data",
			args: args{
				input: dataTxt,
			},
			want: 1909,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfObstructions(tt.args.input); got != tt.want {
				t.Errorf("getNumberOfObstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
