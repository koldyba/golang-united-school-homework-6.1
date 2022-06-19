package golang_united_school_homework

type Shape interface {
	// CalcPerimeter returns calculation result of perimeter
	CalcPerimeter() float64
	// CalcArea returns calculation result of area
	CalcArea() float64
	// GetShapeType returns type of shape
	GetShapeType() string
}
