package e_array

import (
	"strconv"
	"strings"
)

// (1,2,55,999) -> []int{1,2,55,999}
func ArrayI(arg ...int) []int {
	data := make([]int, len(arg))
	for i, val := range arg {
		data[i] = val
	}
	return data
}

// (1.2, 1.3, 0.9) -> []float64{1.2, 1.3, 0.9}
// () -> []float64
func ArrayF(arg ...float64) []float64 {
	data := make([]float64, len(arg))
	for i, val := range arg {
		data[i] = val
	}
	return data
}

// ("1","a", "abcd123") -> []string{"1","a","abcd123"}
func ArrayS(arg ...string) []string {
	data := make([]string, len(arg))
	for i, val := range arg {
		data[i] = val
	}
	return data
}

// ("1,2,3,5, abc, f1f2f3f4") -> []string{"1","2","3","5", "abc", "f1f2f3f4"}
func ArraySS(arg string) []string {
	return strings.Split(arg, ",")
}

// ("1,2,3,5") -> []int{1, 2, 3, 5}
func ArraySI(arg string) []int {
	dt := strings.Split(arg, ",")
	data := make([]int, len(dt))
	for i, val := range dt {
		value, err := strconv.Atoi(strings.TrimSpace(val))
		if err == nil {
			data[i] = value
		}
	}
	return data
}

// ("1.2,1.3") -> []float64{1.2, 1.3}
func ArraySF(arg string) []float64 {
	dt := strings.Split(arg, ",")
	data := make([]float64, len(dt))
	for i, val := range dt {
		value, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err == nil {
			data[i] = value
		}
	}
	return data
}
