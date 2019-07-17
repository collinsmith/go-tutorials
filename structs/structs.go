package structs

import "math"

type Shape interface {
  area() float64
}

type Rect struct {
  width float64
  height float64
}

func perimeter(r Rect) float64 {
  return r.width + r.width + r.height + r.height
}

func (r Rect) area() float64 {
  return r.width * r.height
}

type Circle struct {
  radius float64
}

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

type Triangle struct {
  base float64
  height float64
}

func (t Triangle) area() float64 {
  return 0.5 * t.base * t.height
}