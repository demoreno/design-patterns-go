package main

import "fmt"

// Liskov substitution principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

type Square struct {
	Rectangle
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(size Sized) {
	width := size.GetWidth()
	size.SetHeight(10)
	expectedArea := 10 * width
	actualArea := size.GetWidth() * size.GetHeight()
	fmt.Println(expectedArea, actualArea)
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

type Square2 struct {
	size int // widht and height
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)
	sq := NewSquare(5)
	UseIt(sq)
}
