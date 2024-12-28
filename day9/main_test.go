package main

import (
	_ "embed"
	"testing"
)

//go:embed data.txt
var dataTxt string

func Test_getFsChecksum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getFsChecksum with example data",
			args: args{
				input: "2333133121414131402",
			},
			want: 1928,
		},
		{
			name: "getFsChecksum with task data",
			args: args{
				input: dataTxt,
			},
			want: 6448989155953,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFsChecksum(tt.args.input); got != tt.want {
				t.Errorf("getFsChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCompactFsChecksum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getCompactFsChecksum with example data",
			args: args{
				input: "2333133121414131402",
			},
			want: 2858,
		},
		{
			name: "getCompactFsChecksum with task data",
			args: args{
				input: dataTxt,
			},
			want: 6476642796832,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCompactFsChecksum(tt.args.input); got != tt.want {
				t.Errorf("getCompactFsChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
