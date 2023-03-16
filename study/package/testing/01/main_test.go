package main

import "testing"

//func TestAdd(t *testing.T) {
//	sum := Add(1, 2)
//	if sum == 3 {
//		t.Log("ok")
//	} else {
//		t.Fatal("wrong")
//	}
//}

func TestAdd1(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "add test",
			args: args{
				a: 3,
				b: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
