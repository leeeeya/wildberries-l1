package main

import "testing"

func Test_printRes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printRes(tt.args.s)
		})
	}
}