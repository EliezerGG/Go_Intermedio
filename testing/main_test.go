package main

import "testing"

func TestSum(t *testing.T) {
	//total := Sum(5, 5)
	//if total != 10 {
	//	t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	//}

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", total, item.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{1, 5, 5},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got: %d, want: %d.", max, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		n int
		f int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.n)
		if fib != item.f {
			t.Errorf("Fibonacci was incorrect, got: %d, want: %d.", fib, item.f)
		}
	}
}
