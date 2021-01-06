package main

import (
	"fmt"
)

type figure struct {
	x, y int
}

type values figure

func (f figure) Area() int {
	return (f.x * f.y)
}

func (f values) ValuesType() values {
	return f
}

func (f *figure) Inc(m int) {
	f.x = f.x + m
	f.y = f.y + m
}

func main() {
	var f = figure{x: 3, y: 4}
	f.Inc(10)

	fmt.Println("Area:", f.Area())
	fmt.Println("values:", values(f).ValuesType())
}
