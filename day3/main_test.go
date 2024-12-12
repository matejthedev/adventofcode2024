package main

import "testing"

func Test_calculateResult(t *testing.T) {
	type args struct {
		corruptedData string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test example data",
			args: args{
				corruptedData: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			want: 161,
		},
		{
			name: "real data test",
			args: args{
				corruptedData: dataTxt,
			},
			want: 187833789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateResult(tt.args.corruptedData); got != tt.want {
				t.Errorf("calculateResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateConditionalResult(t *testing.T) {
	type args struct {
		corruptedData string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "PART TWO: test example data",
			args: args{
				corruptedData: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: 48,
		},
		{
			name: "PART TWO: real data test",
			args: args{
				corruptedData: dataTxt,
			},
			want: 94455185,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateConditionalResult(tt.args.corruptedData); got != tt.want {
				t.Errorf("calculateConditionalResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
