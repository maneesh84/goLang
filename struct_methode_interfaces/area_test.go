package structmethodeinterfaces

import (
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("area for rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 6.0}
		got := rectangle.area()
		want := 30.0

		if got != want {
			t.Errorf("got %.2f but want %.2f ", got, want)
		}
	})
	t.Run("area for circle", func(t *testing.T) {
		circle := Circle{5.0}
		got := circle.area()
		want := 31.400000000000002 

		if got != want {
			t.Errorf("got %g but want %g ", got, want)
		}
	})
}