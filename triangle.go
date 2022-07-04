package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (triangle Triangle) CalcPerimeter() float64 {
	return triangle.Side * 3
}

func (triangle Triangle) CalcArea() float64 {
	return (math.Pow(triangle.Side, 2) * math.Sqrt(3)) / 4
}
