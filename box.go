package golang_united_school_homework

import (
	"errors"
	"reflect"
)

const TNSI = "There is no such index"

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

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("Box is full")
	} else {
		b.shapes = append(b.shapes, shape)
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errors.New(TNSI)
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var oldShape Shape
	if i < 0 || i >= len(b.shapes) {
		return nil, errors.New(TNSI)
	}
	oldShape = b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return oldShape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var oldShape Shape
	if i < 0 || i >= len(b.shapes) {
		return nil, errors.New(TNSI)
	}
	oldShape = b.shapes[i]
	b.shapes[i] = shape
	return oldShape, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (sum float64) {
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (sum float64) {
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	// var circleCount int
	// newB := []Shape{}
	// for _, v := range b.shapes {
	// 	_, ok := v.(Circle)
	// 	if !ok {
	// 		newB = append(newB, v)
	// 	} else {
	// 		circleCount++
	// 	}
	// }
	// if circleCount == 0 {
	// 	return errors.New("There are no circles")
	// }
	// b.shapes = newB
	// return nil
	circle := &Circle{}
	c := 0
	var err error
	for i := 0; i < len(b.shapes); i++ {
		if reflect.TypeOf(b.shapes[i]) == reflect.TypeOf(circle) {
			copy(b.shapes[i:], b.shapes[i+1:])
			b.shapes = b.shapes[:len(b.shapes)-1]
			c = c + 1
			i = i - 1
		}
	}
	if c == 0 {
		err = errors.New("There are no circles")
	} else {
		err = nil
	}
	return err
}
