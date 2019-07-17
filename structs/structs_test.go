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
  rect := Rect{12.0, 6.0}
  got := area(rect)
  want := 72.0

  if got != want {
    t.Errorf("got %.2f want %.2f", got, want)
  }
}