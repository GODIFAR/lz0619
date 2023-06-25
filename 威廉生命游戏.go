package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义列行为常量
const width, height = 85, 15

// universe 二维网格
type Universe [][]bool

// 开天辟地 返回空白新世界
func NewUniverse() Universe {
	u := make(Universe, height) // 定框架，高度为height  , [[] [] [] []]
	for i := range u {
		u[i] = make([]bool, width) // 赋值给每一行新元素，bool默认值为0 , [[false false] [] [] []]
	}
	return u
}

// 适者生存
// 生存还是死亡?
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// 统计邻近细胞
func (u Universe) Neighbore(x, y int) int {
	n := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if u.Alive(x+i, y+j) && !(i == 0 && j == 0) {
				n++
			}
		}
	}
	return n
}

// 游戏逻辑  判断下一代存活或者死亡
func (u Universe) Next(x, y int) bool {
	n := u.Neighbore(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

// 激活细胞
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

// 指定细胞状态
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

// 新的世界  将世界a状态更新在世界b中
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

// 观察世界
func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)
	for y := 0; y < height; y++ { //行
		for x := 0; x < width; x++ { //列
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

// 清屏并打印世界
func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a //交换世界
	}

}
