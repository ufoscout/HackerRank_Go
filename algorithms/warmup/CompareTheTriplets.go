package warmup

import "fmt"

func main3() {
	a := make([]int, 3)
	b := make([]int, 3)

	fmt.Scan(&a[0])
	fmt.Scan(&a[1])
	fmt.Scan(&a[2])

	fmt.Scan(&b[0])
	fmt.Scan(&b[1])
	fmt.Scan(&b[2])

	result := calculateScore(a, b)
	fmt.Printf("%v %v", result[0], result[1])
}

func calculateScore(a []int, b []int) []int {
	result := make([]int, 2)
	compareValues(a[0], b[0], result)
	compareValues(a[1], b[1], result)
	compareValues(a[2], b[2], result)
	return result
}

func compareValues(a int, b int, result []int) {
	if (a > b) {
		result[0] = result[0] + 1
	}
	if (a < b) {
		result[1] = result[1] + 1
	}
}