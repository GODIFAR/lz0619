package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 蜜蜂
type honeyBee struct {
	name string
}

func (hb honeyBee) String() string {
	return hb.name
}

func (hb honeyBee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "buzzes about"
	default:
		return "flies to infinity and beyond"
	}
}

func (hb honeyBee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "pollen"
	default:
		return "nectar"
	}
}

// 小田鼠
type gopher struct {
	name string
}

func (g gopher) String() string {
	return g.name
}

func (gb gopher) move() string {
	switch rand.Intn(2) {
	case 0:
		return "scurries along the ground"
	default:
		return "burrows in the sand"
	}
}

func (g gopher) eat() string {
	switch rand.Intn(5) {
	case 0:
		return "carrot"
	case 1:
		return "lettuce"
	case 2:
		return "radish"
	case 3:
		return "corn"
	default:
		return "root"
	}
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v eat the %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	//随机数种子，每次随机数都不是重复的
	rand.Seed(time.Now().UnixNano())

	animals := []animal{
		honeyBee{name: "Buzz Lightyera"},
		gopher{name: "Go gopher"},
	}
	var sol, hour int
	for {
		//打印时间
		fmt.Printf("%2d:00 ", hour)

		if hour < sunrise || hour >= sunset {
			fmt.Println("The animal are sleeping.")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}
		time.Sleep(500 * time.Millisecond)
		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}

}
