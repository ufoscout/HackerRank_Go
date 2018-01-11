package warmup

import "fmt"

// SolveMeFirst ...
func solveMeFirst(a uint32, b uint32) uint32 {
	return (a + b)
}

func main1() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
