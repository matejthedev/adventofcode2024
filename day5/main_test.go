package main

import (
	_ "embed"
	"testing"
)

//go:embed data.txt
var dataTxt string

func Test_getMiddlePageNumSum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get middle page numbers sum from example data",
			args: args{
				input: "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47",
			},
			want: 143,
		},
		{
			name: "get middle page numbers sum from the task data",
			args: args{
				input: dataTxt,
			},
			want: 4185,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMiddlePageNumSum(tt.args.input); got != tt.want {
				t.Errorf("getMiddlePageNumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getOrderedMiddlePageNumSum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get ordered page numbers sum from the task data",
			args: args{
				input: "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47",
			},
			want: 123,
		},
		{
			name: "get ordered page numbers sum from example data",
			args: args{
				input: dataTxt,
			},
			want: 4480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrderedMiddlePageNumSum(tt.args.input); got != tt.want {
				t.Errorf("getOrderedMiddlePageNumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
