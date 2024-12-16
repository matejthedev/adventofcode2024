package main

import (
	_ "embed"
	"testing"
)

//go:embed data.txt
var dataTxt string

func Test_getXmasCount(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get Xmas count example data",
			args: args{
				input: "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX",
			},
			want: 18,
		},
		{
			name: "get Xmas count from the task data",
			args: args{
				input: dataTxt,
			},
			want: 2573,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getXmasCount(tt.args.input); got != tt.want {
				t.Errorf("getXmasCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countCrossedMasCount(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get crossed word \"mas\" count from example data",
			args: args{
				input: "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX",
			},
			want: 9,
		},
		{
			name: "get crossed word \"mas\" count from the task data",
			args: args{
				input: dataTxt,
			},
			want: 1850,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCrossedMasCount(tt.args.input); got != tt.want {
				t.Errorf("countCrossedMasCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
