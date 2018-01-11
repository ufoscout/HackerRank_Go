package warmup

import "fmt"

func main2() {
	var a, temp, res int
	fmt.Scan(&a)

	for index := 0; index < a; index++ {
		fmt.Scan(&temp)
		res = res + temp
	}

	fmt.Println(res)
}
