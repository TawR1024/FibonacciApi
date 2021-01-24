package calculator

import "testing"

func Test_binet(t *testing.T) {

	tests := []struct {
		name     string
		position int
		want     float64
	}{
		{"Test  position 1", 1, 1},
		{"Test  position 2", 2, 1},
		{"Test  position 11", 11, 89},
		{"Test  position 22", 22, 17711},
		{"Test  position 50", 50, 12586269025},
		{"Test  position 90", 90, 2880067194370816120}, //2.8800671943708155e+18, want 2.880067194370816e+18
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binet(tt.position); got != tt.want {
				t.Errorf("binet() = %v, want %v", got, tt.want)
			}
		})
	}
}
