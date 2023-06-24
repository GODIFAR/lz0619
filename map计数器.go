package main

import (
	"fmt"
)

// 计数器 统计数组中每个数字出现了几次
func main() {
	tempratures := []float64{-28.0, -59.0, -26.0, -66.0, -59.0}

	frequency := make(map[float64]int)

	for _, t := range tempratures {
		frequency[t]++
		fmt.Println(frequency)
	}
	for t, num := range frequency {
		fmt.Println(t, num)
	}

}
