package structmethodeinterfaces

import "math"

type Rectangle struct{
	height float64
	width float64
}
func (r Rectangle)area() float64{
	return r.height*r.width
}


type Circle struct{
	radius float64
}
func (c Circle) area() float64{
	return math.Pi *(c.radius*c.radius)
}

type Triangle struct{
	height float64
	base float64
}
func (t Triangle) area() float64 {
	return t.height*t.base*0.5
}

type Shape interface{
	area() float64
}