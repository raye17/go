package main

import "testing"

func Test_swap(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "001-swap", args: args{
			a: 12,
			b: 18,
		}},
		{name: "002", args: args{
			a: 21,
			b: 23,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.a, tt.args.b)
		})
	}
}

func Test_xor(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "003-xor", args: args{
			a: 33,
			b: 44,
		}},
		{name: "004", args: args{
			a: 55,
			b: 66,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xor(tt.args.a, tt.args.b)
		})
	}
}

func Benchmark_swap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		swap(3, 6)
	}
}

func Benchmark_xor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xor(4, 5)
	}
}
