package main

func SumAll(numA []int, numB []int) []int {
	sumA := 0
	sumB := 0
	for _, number := range numA {
		sumA += number
	}
	for _, number := range numB {
		sumB += number
	}

	return []int{sumA, sumB}
}
