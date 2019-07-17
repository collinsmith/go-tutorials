package structs

type Rect struct {
  width float64
  height float64
}

func perimeter(r Rect) float64 {
  return r.width + r.width + r.height + r.height
}

func area(r Rect) float64 {
  return r.width * r.height
}