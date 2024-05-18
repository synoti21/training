package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area()
}

type Circle struct {
	rad float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c *Circle) Area() float64 {
	return math.Pow(c.rad, 2) * 3.14
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func NewCircle(r float64) *Circle {
	return &Circle{
		rad: r,
	}
}

func NewRectangle(w float64, h float64) *Rectangle {
	return &Rectangle{
		width:  w,
		height: h,
	}
}

func getArea(s Shape) {
	s.Area()
}

func main() {
	circle := NewCircle(0.0)
	rec := NewRectangle(0.0, 0.0)

	fmt.Printf("원의 반지름을 입력하세요 :")
	fmt.Scanf("%f", &circle.rad)
	fmt.Printf("사각형의 가로를 입력하세요 :")
	fmt.Scanf("%f", &rec.height)
	fmt.Printf("사각형의 세로를 입력하세요 :")
	fmt.Scanf("%f", &rec.width)

	fmt.Printf("원의 면적: %f", circle.Area())
	fmt.Printf("사각형의 면적: %f", rec.Area())
}

