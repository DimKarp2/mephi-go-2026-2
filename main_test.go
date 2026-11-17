package main

import (
	"math/rand"
	"testing"
)

// TestGenerateRandomElements проверяет генерацию случайных чисел,
// корректность размера слайса и обработку неположительных размеров.
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "negative size", size: -5},
		{name: "zero size", size: 0},
		{name: "one element", size: 1},
		{name: "many elements", size: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != max(tt.size, 0) {
				t.Fatalf("got slice length %d, want %d", len(data), max(tt.size, 0))
			}

			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive integer", i, value)
				}
			}
		})
	}
}

// TestMaximum проверяет поиск максимального значения
// для различных вариантов входных данных,
// включая пустые слайсы, отрицательные числа и повторяющиеся значения.
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "maximum in the middle", data: []int{3, 15, 8}, want: 15},
		{name: "descending values", data: []int{9, 7, 4, 1}, want: 9},
		{name: "negative values", data: []int{-8, -2, -10}, want: -2},
		{name: "mixed values", data: []int{-5, 0, 12, 4}, want: 12},
		{name: "equal values", data: []int{6, 6, 6}, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

// TestMaxChunks проверяет параллельный поиск максимума,
// включая случаи, когда количество элементов меньше количества частей,
// а также когда длина слайса не делится на количество частей без остатка.
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "fewer elements than chunks", data: []int{2, 9, 4}, want: 9},
		{name: "exactly eight elements", data: []int{1, 2, 3, 4, 5, 6, 7, 20}, want: 20},
		{name: "length not divisible by chunks", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99, 10}, want: 99},
		{name: "maximum at the end", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}, want: 100},
		{name: "negative values", data: []int{-8, -2, -10, -4, -5, -6, -7, -9, -3}, want: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

// TestMaxChunksMatchesMaximum проверяет, что результаты
// последовательного и параллельного поиска максимума совпадают
// для слайсов различной длины со случайными значениями.
func TestMaxChunksMatchesMaximum(t *testing.T) {
	rng := rand.New(rand.NewSource(1))

	for size := 0; size <= 100; size++ {
		data := make([]int, size)
		for i := range data {
			data[i] = rng.Intn(2_001) - 1_000
		}

		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("size %d: maxChunks() = %d, maximum() = %d", size, got, want)
		}
	}
}
