package calculator

import (
	"math/big"
	"testing"
)

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

func Test_fibonacciBig(t *testing.T) {
	f1 := big.NewInt(1)
	f2 := big.NewInt(1)

	f44 := big.NewInt(1)
	f44.SetString("701408733", 10)

	f90 := big.NewInt(1)
	f90.SetString("2880067194370816120", 10)

	f124 := big.NewInt(1)
	f124.SetString("36726740705505779255899443", 10)

	f0 := big.NewInt(10)
	tests := []struct {
		name     string
		position int
		want     *big.Int
	}{
		// use info from wolfram alpha
		{"Test  position 1", 1, f1},
		{"Test  position 2", 2, f2},
		{"Test  position 44", 44, f44},
		{"Test  position 90", 90, f90},
		{"Test  position 124", 124, f124},
		{"Test  position 0", 0, f0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibonacciBig(tt.position)
			if got.String() != tt.want.String() {
				t.Errorf("fibonacciBig() = %v, want %v", got, tt.want)
			}
		})
	}
}
