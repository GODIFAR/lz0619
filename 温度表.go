package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64

func (c celsius) cTof() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}
func (f fahrenheit) fToc() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "==================="
	rowFormat    = "|%8s|%8s|\n"
	numberFormat = "%.1f"
)

type getRowFormat func(row int) (string, string)

// 画出表格
func drawTable(h1, h2 string, rows int, getrow getRowFormat) {
	fmt.Println(line)
	fmt.Printf(rowFormat, h1, h2)
	fmt.Println(line)

	for row := 0; row < rows; row++ {

		c1, c2 := getrow(row)
		fmt.Printf(rowFormat, c1, c2)
	}
}

func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.cTof()
	cel1 := fmt.Sprintf(numberFormat, c)
	cel2 := fmt.Sprintf(numberFormat, f)
	return cel1, cel2
}

func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.fToc()
	cel1 := fmt.Sprintf(numberFormat, f)
	cel2 := fmt.Sprintf(numberFormat, c)
	return cel1, cel2
}

func main() {
	drawTable(".C", ".F", 29, ctof)
	fmt.Println()
	drawTable(".F", ".C", 29, ftoc)
}
