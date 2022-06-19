package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

//errors list
var (
	errorOutOfCap  = errors.New("no space left in the box")
	errorBadIndex  = errors.New("incorrect index")
	errorNoCircles = errors.New("there are no circles in the box")
)

func (b *box) checkBoxIndex(i int) error {
	if i >= b.shapesCapacity || b.shapes[i] == nil {
		return errorBadIndex
	}
	return nil
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errorOutOfCap
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	err := b.checkBoxIndex(i)
	if err != nil {
		return nil, err
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	b.shapes[i] = b.shapes[len(b.shapes)-1]
	b.shapes = b.shapes[:len(b.shapes)-1]
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	b.shapes[i] = shape
	return s, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for i := range b.shapes {
		s, err := b.GetByIndex(i)
		if err != nil {
			continue
		}
		sum += s.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for i := range b.shapes {
		s, err := b.GetByIndex(i)
		if err != nil {
			continue
		}
		sum += s.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	err := errorNoCircles
	i := 0
	for _, s := range b.shapes {
		if s.GetShapeType() != "circle" {
			b.shapes[i] = s
			i++
		} else {
			err = nil
		}
	}
	b.shapes = b.shapes[:i]
	return err
}
