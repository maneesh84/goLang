package structmethodeinterfaces

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
	return 2*3.14*c.radius
}