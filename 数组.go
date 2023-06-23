package main

import (
	"fmt"
)

// 值传递接受参数，长度不符的数组作为参数传递会报错
func terraform(planets [3]string) {
	for i := range planets {
		planets[i] = "new " + planets[i]
	}
	fmt.Println(planets)
}
func main() {
	//var planets [8]string
	planets := [...]string{
		"a",
		"b",
		"c",
	}
	/*range 遍历数组 可以避免数组越界
	  for适合指定数量数据使用，或者指定索引*/

	for i, planet := range planets {
		fmt.Println(i, planet)
	}
	// planet 输出数据的第i个值，planets输出的是数组
	for i := range planets {
		fmt.Println(i, planets)
	}

	for j := 0; j < len(planets); j++ {
		fmt.Println(j, planets[j])
	}
	//数组值传递
	p := planets
	planets[2] = "k"
	fmt.Println(p)
	fmt.Println(planets)

	terraform(planets)

	//二维数组
	var jz [6][6]string
	jz[0][0] = "r"
	jz[5][5] = "r"
	for i := range jz {
		jz[1][i] = "p"
	}
	fmt.Println(jz)
}
