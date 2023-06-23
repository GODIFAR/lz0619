package main

import (
	"fmt"
	"strings"
)

func main() {

	x := " sddsdsd sdsdhgkf hsdjgkhf jkg "
	key := "golang"
	a := ""
	b := 0
	x = strings.ToUpper(strings.Replace(x, " ", "", -1))
	key = strings.ToUpper(key)
	for i := 0; i < len(x); i++ {
		c := x[i]
		if c >= 'A' && c <= 'Z' {
			c -= 'A'
			d := key[b] - 'A'
			c = (c+d)%26 + 'A'
			b++
			b %= len(key)
		}

		a = a + string(c)
	}
	fmt.Println(a)
}
