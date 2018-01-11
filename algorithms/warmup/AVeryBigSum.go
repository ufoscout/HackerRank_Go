package warmup

import "fmt"

func main4() {
	var n int
	fmt.Scan(&n)
	var values = make([]int64, n)

	for index := 0; index < n; index++ {
		fmt.Scan(&values[index])
	}

	fmt.Print(sumAVeryBigSum(values))
}

func sumAVeryBigSum(values []int64) int64 {
	var result int64
	for _, value := range values {
		result = result + value
	}
	return result
}
