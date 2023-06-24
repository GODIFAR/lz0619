package main

import (
	"fmt"
)

func main() {
	//切片
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	//切分数组
	b := a[:4]
	c := a[4:6]
	d := a[6:]
	fmt.Println(a, b, c, d)
	//切分字符串
	r1 := "abcdefgh"
	tune := r1[4:]
	fmt.Println(tune)
	//修改字符串不会改变切片的值
	r1 = "abch"
	fmt.Println(tune)

	// 切片 Slice 切分数组声明
	dwarfArry := [...]string{"a", "b", "c", "d"}
	dwarf := dwarfArry[:]
	// 切片 Slice 直接声明
	dwarfarry := []string{"a", "s", "d", "f"}
	fmt.Printf("%T,%T", dwarf, dwarfarry)
	fmt.Println(dwarf, dwarfarry)
}
