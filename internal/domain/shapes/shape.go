package shapes

import (
	"errors"

	"github.com/rs/xid"
)

type ShapeRepository interface {
	Save(shape *ShapeObject) error
	GetByID(id string) (*ShapeObject, error)
}

type IShape interface {
	GetParameters() ShapeParameters
	GetArea() float64
	GetPerimeter() float64
}

type Pointer struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ShapeParameters map[string]int

type ShapeObject struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	ShapeType   string  `json:"shapeType"`
	ActualPoint Pointer `json:"actualPoint"`
	Shape       any     `json:"shape"`
}

func NewShapeObject(description, shapeType string, startPoint Pointer, params ShapeParameters) (*ShapeObject, error) {
	shapeFact, ok := shapeFactory[shapeType]
	if !ok {
		return nil, errors.New("invalid shape type")
	}
	shape, err := shapeFact(params)
	if err != nil {
		return nil, err
	}

	return &ShapeObject{
		ID:          xid.New().String(),
		Description: description,
		ShapeType:   shapeType,
		ActualPoint: startPoint,
		Shape:       shape,
	}, nil
}

func (s *ShapeObject) getIShape() IShape {
	switch s.ShapeType {
	case "circle":
		return s.Shape.(*Circle)
	case "square":
		return s.Shape.(*Square)
	case "rectangle":
		return s.Shape.(*Rectangle)
	case "triangle":
		return s.Shape.(*Triangle)
	default:
		return nil
	}
}

func (s *ShapeObject) GetParameters() ShapeParameters {
	return s.getIShape().GetParameters()
}

func (s *ShapeObject) GetArea() float64 {
	return s.getIShape().GetArea()
}

func (s *ShapeObject) GetPerimeter() float64 {
	return s.getIShape().GetPerimeter()
}

/*****************************************************************
   Here we created objects that implement the IShape interface
******************************************************************/

var shapeFactory = map[string]func(ShapeParameters) (any, error){
	"circle":    newCircle,
	"square":    newSquare,
	"rectangle": newRectangle,
	"triangle":  newTriangle,
}

// Circle is a simple representation of a circle shape.
type Circle struct {
	Radius int `json:"radius"`
}

func newCircle(params ShapeParameters) (any, error) {
	if radius, ok := params["radius"]; ok {
		return &Circle{Radius: radius}, nil
	}
	return nil, errors.New("invalid parameters: Circle requires a 'radius' parameter")
}

func (c Circle) GetParameters() ShapeParameters {
	return ShapeParameters{"radius": c.Radius}
}

func (c Circle) GetArea() float64 {
	return 3.14 * float64(c.Radius*c.Radius)
}

func (c Circle) GetPerimeter() float64 {
	return 2 * 3.14 * float64(c.Radius)
}

// Square is a simple representation of a square shape.
type Square struct {
	Side int `json:"side"`
}

func newSquare(params ShapeParameters) (any, error) {
	if side, ok := params["side"]; ok {
		return &Square{Side: side}, nil
	}
	return nil, errors.New("invalid parameters: Square requires a 'side' parameter")
}

func (s Square) GetParameters() ShapeParameters {
	return ShapeParameters{"side": s.Side}
}

func (s Square) GetArea() float64 {
	return float64(s.Side * s.Side)
}

func (s Square) GetPerimeter() float64 {
	return float64(4 * s.Side)
}

// Rectangle is a simple representation of a rectangle shape.
type Rectangle struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func newRectangle(params ShapeParameters) (any, error) {
	if width, ok := params["width"]; ok {
		if height, ok := params["height"]; ok {
			return &Rectangle{Width: width, Height: height}, nil
		}
		return nil, errors.New("invalid parameters: Rectangle requires a 'height' parameter")
	}
	return nil, errors.New("invalid parameters: Rectangle requires a 'width' parameter")
}

func (r Rectangle) GetParameters() ShapeParameters {
	return ShapeParameters{"width": r.Width, "height": r.Height}
}

func (r Rectangle) GetArea() float64 {
	return float64(r.Width * r.Height)
}

func (r Rectangle) GetPerimeter() float64 {
	return float64(2 * (r.Width + r.Height))
}

// Triangle is a simple representation of a triangle shape.
type Triangle struct {
	Base   int `json:"base"`
	Height int `json:"height"`
}

func newTriangle(params ShapeParameters) (any, error) {
	if base, ok := params["base"]; ok {
		if height, ok := params["height"]; ok {
			return &Triangle{Base: base, Height: height}, nil
		}
		return nil, errors.New("invalid parameters: Triangle requires a 'height' parameter")
	}
	return nil, errors.New("invalid parameters: Triangle requires a 'base' parameter")
}

func (t Triangle) GetParameters() ShapeParameters {
	return ShapeParameters{"base": t.Base, "height": t.Height}
}

func (t Triangle) GetArea() float64 {
	return 0.5 * float64(t.Base*t.Height)
}

func (t Triangle) GetPerimeter() float64 {
	// Assuming an equilateral triangle for simplicity
	return float64(3 * t.Base)
}
