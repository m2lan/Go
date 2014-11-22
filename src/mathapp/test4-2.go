package main

import (
	"fmt"
)

/* 自定义类型中定义任意多的method */

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// 定义Color作为byte的别名
type Color byte

// 定义一个struct:Box,含有三个长宽高字段和一个颜色属性
type Box struct {
	width, height, depth float64
	color                Color
}

// 定义一个slice:BoxList,含有Box
type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) setColor(c Color) {
	b.color = c
}

func (b1 BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}

	return k
}

func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
		b1[i].setColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("xxxx---%d\n", len(boxes))
	fmt.Println(boxes[0].Volume())
	fmt.Println(boxes[len(boxes)-1].color.String())
	fmt.Println(boxes.BiggestsColor().String())

	boxes.PaintItBlack()
	fmt.Println(boxes[1].color.String())
	fmt.Println(boxes.BiggestsColor().String())
}
