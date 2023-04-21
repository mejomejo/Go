package main

import "fmt"

type Point struct {
	X, Y int
}

type Shape interface {
	Draw()
	Move(dx, dy int)
}

type Rect struct {
	LeftTop, RightBottom Point
}

func (r *Rect) Draw() {
	fmt.Printf("Drawing Rectangle with LeftTop(%d, %d) and RightBottom(%d, %d)\n",
		r.LeftTop.X, r.LeftTop.Y, r.RightBottom.X, r.RightBottom.Y)
}

func (r *Rect) Move(dx, dy int) {
	r.LeftTop.X += dx
	r.LeftTop.Y += dy
	r.RightBottom.X += dx
	r.RightBottom.Y += dy
}

type Circle struct {
	Center Point
	Radius int
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing Circle with Center(%d, %d) and Radius(%d)\n",
		c.Center.X, c.Center.Y, c.Radius)
}

func (c *Circle) Move(dx, dy int) {
	c.Center.X += dx
	c.Center.Y += dy
}

type Triangle struct {
	Point1, Point2, Point3 Point
}

func (t *Triangle) Draw() {
	fmt.Printf("Drawing Triangle with Point1(%d, %d), Point2(%d, %d) and Point3(%d, %d)\n",
		t.Point1.X, t.Point1.Y, t.Point2.X, t.Point2.Y, t.Point3.X, t.Point3.Y)
}

func (t *Triangle) Move(dx, dy int) {
	t.Point1.X += dx
	t.Point1.Y += dy
	t.Point2.X += dx
	t.Point2.Y += dy
	t.Point3.X += dx
	t.Point3.Y += dy
}

func main() {
	shapes := []Shape{
		&Rect{Point{0, 0}, Point{100, 100}},
		&Circle{Point{50, 50}, 25},
		&Triangle{Point{0, 0}, Point{50, 50}, Point{100, 0}},
	}

	for _, shape := range shapes {
		shape.Draw()
		shape.Move(10, 10)
		shape.Draw()
	}
}
