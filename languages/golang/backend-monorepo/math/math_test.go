package math

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 2, 3, 5},
		{"zero values", 0, 0, 0},
		{"negative numbers", -1, -2, -3},
		{"mixed signs", -5, 10, 5},
		{"large numbers", 1_000_000, 2_000_000, 3_000_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Math{A: tt.a, B: tt.b}
			if got := m.Add(); got != tt.want {
				t.Errorf("Add() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	m := Math{A: 100, B: 200}
	for range b.N {
		m.Add()
	}
}
