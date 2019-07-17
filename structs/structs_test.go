package structs

import "testing"

func TestPerimeter(t *testing.T) {
  rect := Rect{10.0, 10.0}
  got := perimeter(rect)
  want := 40.0

  if got != want {
    t.Errorf("got %.2f want %.2f", got, want)
  }
}

func TestArea(t *testing.T) {
  areaTests := []struct {
    shape Shape
    want  float64
  }{
    {Rect{12, 6}, 72.0},
    {Circle{10}, 314.1592653589793},
    {Triangle{12, 6}, 36.0},
  }

  for _, tt := range areaTests {
    got := tt.shape.area()
    if got != tt.want {
      t.Errorf("got %.2f want %.2f", got, tt.want)
    }
  }
}