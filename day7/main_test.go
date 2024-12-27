package main

import (
	_ "embed"
	"testing"
)

//go:embed data.txt
var dataTxt string

func Test_getTotalCalibrationResult(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getTotalCalibrationResult with example data",
			args: args{
				input: "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20",
			},
			want: 3749,
		},
		{
			name: "getTotalCalibrationResult with task data",
			args: args{
				input: dataTxt,
			},
			want: 6083020304036,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalCalibrationResult(tt.args.input); got != tt.want {
				t.Errorf("getTotalCalibrationResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTotalResultWithInsertingOperators(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getTotalResultWithInsertingOperators with example data",
			args: args{
				input: "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20",
			},
			want: 11387,
		},
		{
			name: "getTotalResultWithInsertingOperators with task data",
			args: args{
				input: dataTxt,
			},
			want: 59002246504791,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalResultWithInsertingOperators(tt.args.input); got != tt.want {
				t.Errorf("getTotalResultWithInsertingOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}
