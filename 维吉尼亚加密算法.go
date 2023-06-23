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
	//strings.ToUpper  小写字母转换成大写字母
	//strings.Replace  替代部分原字符串
	x = strings.ToUpper(strings.Replace(x, " ", "", -1))
	key = strings.ToUpper(key)
	for i := 0; i < len(x); i++ {
		c := x[i]
		if c >= 'A' && c <= 'Z' {
			c -= 'A'
			d := key[b] - 'A'
			//加密字母+关键字字母
			c = (c+d)%26 + 'A'
			b++
			b %= len(key)
		}

		a = a + string(c)
	}
	fmt.Println(a)
}
