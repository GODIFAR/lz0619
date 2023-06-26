package main

import (
	"fmt"
)

// location结构包含精度维度
type location struct {
	lat, long float64
}

// coordinate 结构使用度/分/秒（dms)格式的坐标表示东西南北半球
type coordinate struct {
	d, m, s float64
	h       rune
}

// newLocation函数接收DMS格式的维度和经度坐标，创建出相应的十进制位置
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

// decimal 方法会将DMS格式的坐标转换为十进制格式
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'W', 'S', 'w', 's':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
func main() {
	a := newLocation(coordinate{14, 34, 6.2, 's'}, coordinate{175, 28, 21.5, 'E'})
	b := newLocation(coordinate{14, 34, 6.2, 's'}, coordinate{175, 28, 21.5, 'E'})
	c := newLocation(coordinate{14, 34, 6.2, 'n'}, coordinate{175, 28, 21.5, 'E'})
	i := newLocation(coordinate{14, 34, 6.2, 'n'}, coordinate{175, 28, 21.5, 'E'})
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("i", i)
}
