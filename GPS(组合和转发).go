package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64 //半径
}
type location struct {
	name      string
	lat, long float64
}

// 返回包含名称、维度、经度的字符串
func (l location) description() string {
	return fmt.Sprintf("%v(%.1f。.%.1f。)", l.name, l.lat, l.long)
}

type gps struct {
	world       world
	current     location //当前位置
	destination location //目标位置
}

// 计算当前位置和目标位置的距离
func (g gps) distance() float64 {
	return g.world.distances(g.current, g.destination)
}

// 计算当前位置和目标位置的距离
func (w world) distances(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// 输出距离和目标位置的名称、维度、经度
func (g gps) message() string {
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.destination.description())
}

type rover struct {
	gps
}

func main() {
	mars := world{
		radius: 3389.5,
	}
	bradbury := location{
		"Bradbury Landing", -4.585, -137.4417,
	}
	elysium := location{
		"Elysium Planitia", 4.5, 135.9,
	}
	gps := gps{
		world:       mars,
		current:     bradbury,
		destination: elysium,
	}
	curiosty := rover{
		gps: gps,
	}
	fmt.Println(curiosty.message())
}
