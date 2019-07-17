package structs

import "math"

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