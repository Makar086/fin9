package main

import (
	"math/rand"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	rand.Seed(42)

	tests := []struct {
		name         string
		size         int
		wantLen      int
		wantCap      int
		expectEmpty bool
	}{
		{
			name:        "Отрицательный размер",
			size:        -5,
			wantLen:     0,
			wantCap:     0,
			expectEmpty: true,
		},
		{
			name:        "Нулевой размер (пустой слайс)",
			size:        0,
			wantLen:     0,
			wantCap:     0,
			expectEmpty: true,
		},
		{
			name:        "Один элемент",
			size:        1,
			wantLen:     1,
			wantCap:     1,
			expectEmpty: false,
		},
		{
			name:        "Стандартный размер (несколько элементов)",
			size:        10,
			wantLen:     10,
			wantCap:     10,
			expectEmpty: false,
		},
		{
			name:        "Большой размер",
			size:        1000,
			wantLen:     1000,
			wantCap:     1000,
			expectEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			if len(result) != tt.wantLen {
				t.Errorf("generateRandomElements(%d) длина = %d; получили %d", tt.size, tt.wantLen, len(result))
			}

			if cap(result) != tt.wantCap {
				t.Errorf("generateRandomElements(%d) емкость = %d; получили %d", tt.size, tt.wantCap, cap(result))
			}
			
			if tt.expectEmpty && result == nil {
				t.Errorf("generateRandomElements(%d) вернул nil, ожидался инициализированный пустой слайс []int{}", tt.size)
			}

			
			if !tt.expectEmpty {
				for i, val := range result {
					if val < 0 {
						t.Errorf("В индексе %d обнаружено отрицательное число: %d", i, val)
					}
				}
			}
		})
	}
}

func TestGenerateRandomElements_Randomness(t *testing.T) {
	size := 5
	slice1 := generateRandomElements(size)
	slice2 := generateRandomElements(size)

	if len(slice1) != len(slice2) {
		t.Fatalf("Размеры слайсов не совпадают")
	}

	identical := true
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			identical = false
			break
		}
	}

	if identical {
		t.Error("generateRandomElements вернула два идентичных слайса подряд. Убедитесь, что генератор инициализирован случайно.")
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "Empty slice",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "Nil slice",
			data:     nil,
			expected: 0,
		},
		{
			name:     "Single element",
			data:     []int{5},
			expected: 5,
		},
		{
			name:     "All positive numbers",
			data:     []int{1, 3, 9, 4, 7},
			expected: 9,
		},
		{
			name:     "All negative numbers",
			data:     []int{-5, -1, -10, -3},
			expected: -1,
		},
		{
			name:     "Mixed numbers",
			data:     []int{-10, 0, 15, -3, 2},
			expected: 15,
		},
		{
			name:     "Duplicate maximum values",
			data:     []int{8, 2, 8, 5, 8},
			expected: 8,
		},
		{
			name:     "Identical elements",
			data:     []int{4, 4, 4, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maximum(tt.data)
			if actual != tt.expected {
				t.Errorf("maximum(%v) = %d; expected %d", tt.data, actual, tt.expected)
			}
		})
	}
}