package structmethodeinterfaces

import (
	"testing"
)

func TestArea(t *testing.T) {

	// checkSum:= func(t testing.TB, shape Shape, want float64){
	// 	t.Helper()
	// 	got:= shape.area()
	// 	if(got!=want){
	// 		t.Errorf("got %g but want %g", got ,want)
	// 	}
	// }
	// t.Run("area for rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{5.0, 6.0}
	// 	checkSum(t,rectangle,30)
	// })
	// t.Run("area for circle", func(t *testing.T) {
	// 	circle := Circle{5.0}
	// 	checkSum(t,circle,31.400000000000002)
	// })

	testSlice:= []struct{
		shape Shape
		want float64
	}{
		{shape: Rectangle{width: 12, height: 6}, want: 72.0},
        {shape: Circle{radius: 10}, want: 314.1592653589793},
        {shape: Triangle{base: 12, height: 6}, want: 36.0},
	}

	for _,tt:=range testSlice{
		got:=tt.shape.area()
		if(got!=tt.want){
			t.Errorf("%#v got %g but want %g",tt, got ,tt.want)
		}
	}
}