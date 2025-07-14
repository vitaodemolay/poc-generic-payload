package shapes

import (
	"testing"
)

func TestNewCircle(t *testing.T) {
	params := ShapeParameters{"radius": 5}
	c, err := newCircle(params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	circle, ok := c.(*Circle)
	if !ok {
		t.Fatalf("expected *Circle, got %T", c)
	}
	if circle.Radius != 5 {
		t.Errorf("expected radius 5, got %d", circle.Radius)
	}
}

func TestNewCircle_MissingParam(t *testing.T) {
	params := ShapeParameters{}
	_, err := newCircle(params)
	if err == nil {
		t.Fatal("expected error for missing radius, got nil")
	}
}

func TestCircleMethods(t *testing.T) {
	c := Circle{Radius: 3}
	if area := c.GetArea(); area != 28.26 {
		t.Errorf("expected area 28.26, got %v", area)
	}
	if perim := c.GetPerimeter(); perim != 18.84 {
		t.Errorf("expected perimeter 18.84, got %v", perim)
	}
	params := c.GetParameters()
	if params["radius"] != 3 {
		t.Errorf("expected radius param 3, got %d", params["radius"])
	}
}

func TestNewSquare(t *testing.T) {
	params := ShapeParameters{"side": 4}
	s, err := newSquare(params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	square, ok := s.(*Square)
	if !ok {
		t.Fatalf("expected *Square, got %T", s)
	}
	if square.Side != 4 {
		t.Errorf("expected side 4, got %d", square.Side)
	}
}

func TestNewSquare_MissingParam(t *testing.T) {
	params := ShapeParameters{}
	_, err := newSquare(params)
	if err == nil {
		t.Fatal("expected error for missing side, got nil")
	}
}

func TestSquareMethods(t *testing.T) {
	s := Square{Side: 2}
	if area := s.GetArea(); area != 4 {
		t.Errorf("expected area 4, got %v", area)
	}
	if perim := s.GetPerimeter(); perim != 8 {
		t.Errorf("expected perimeter 8, got %v", perim)
	}
	params := s.GetParameters()
	if params["side"] != 2 {
		t.Errorf("expected side param 2, got %d", params["side"])
	}
}

func TestNewRectangle(t *testing.T) {
	params := ShapeParameters{"width": 3, "height": 6}
	r, err := newRectangle(params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	rect, ok := r.(*Rectangle)
	if !ok {
		t.Fatalf("expected *Rectangle, got %T", r)
	}
	if rect.Width != 3 || rect.Height != 6 {
		t.Errorf("expected width 3 and height 6, got %d and %d", rect.Width, rect.Height)
	}
}

func TestNewRectangle_MissingParam(t *testing.T) {
	params := ShapeParameters{"width": 3}
	_, err := newRectangle(params)
	if err == nil {
		t.Fatal("expected error for missing height, got nil")
	}
	params = ShapeParameters{"height": 3}
	_, err = newRectangle(params)
	if err == nil {
		t.Fatal("expected error for missing width, got nil")
	}
}

func TestRectangleMethods(t *testing.T) {
	r := Rectangle{Width: 2, Height: 5}
	if area := r.GetArea(); area != 10 {
		t.Errorf("expected area 10, got %v", area)
	}
	if perim := r.GetPerimeter(); perim != 14 {
		t.Errorf("expected perimeter 14, got %v", perim)
	}
	params := r.GetParameters()
	if params["width"] != 2 || params["height"] != 5 {
		t.Errorf("expected width 2 and height 5, got %d and %d", params["width"], params["height"])
	}
}

func TestNewTriangle(t *testing.T) {
	params := ShapeParameters{"base": 4, "height": 3}
	tr, err := newTriangle(params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	triangle, ok := tr.(*Triangle)
	if !ok {
		t.Fatalf("expected *Triangle, got %T", tr)
	}
	if triangle.Base != 4 || triangle.Height != 3 {
		t.Errorf("expected base 4 and height 3, got %d and %d", triangle.Base, triangle.Height)
	}
}

func TestNewTriangle_MissingParam(t *testing.T) {
	params := ShapeParameters{"base": 4}
	_, err := newTriangle(params)
	if err == nil {
		t.Fatal("expected error for missing height, got nil")
	}
	params = ShapeParameters{"height": 4}
	_, err = newTriangle(params)
	if err == nil {
		t.Fatal("expected error for missing base, got nil")
	}
}

func TestTriangleMethods(t *testing.T) {
	tg := Triangle{Base: 6, Height: 2}
	if area := tg.GetArea(); area != 6 {
		t.Errorf("expected area 6, got %v", area)
	}
	if perim := tg.GetPerimeter(); perim != 18 {
		t.Errorf("expected perimeter 18, got %v", perim)
	}
	params := tg.GetParameters()
	if params["base"] != 6 || params["height"] != 2 {
		t.Errorf("expected base 6 and height 2, got %d and %d", params["base"], params["height"])
	}
}

func TestNewShapeObject_Valid(t *testing.T) {
	params := ShapeParameters{"radius": 7}
	start := Pointer{X: 1, Y: 2}
	obj, err := NewShapeObject("A circle", "circle", start, params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if obj.Description != "A circle" {
		t.Errorf("expected description 'A circle', got %s", obj.Description)
	}
	if obj.ShapeType != "circle" {
		t.Errorf("expected shapeType 'circle', got %s", obj.ShapeType)
	}
	if obj.ActualPoint != start {
		t.Errorf("expected actualPoint %+v, got %+v", start, obj.ActualPoint)
	}
	if obj.Shape == nil {
		t.Error("expected shape to be non-nil")
	}
}

func TestNewShapeObject_InvalidType(t *testing.T) {
	params := ShapeParameters{}
	start := Pointer{X: 0, Y: 0}
	_, err := NewShapeObject("Unknown", "hexagon", start, params)
	if err == nil {
		t.Fatal("expected error for invalid shape type, got nil")
	}
}
func TestShapeObject_getIShape_Circle(t *testing.T) {
	params := ShapeParameters{"radius": 2}
	obj, err := NewShapeObject("circle", "circle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ishape := obj.getIShape()
	circle, ok := ishape.(*Circle)
	if !ok {
		t.Fatalf("expected *Circle, got %T", ishape)
	}
	if circle.Radius != 2 {
		t.Errorf("expected radius 2, got %d", circle.Radius)
	}
}

func TestShapeObject_getIShape_Square(t *testing.T) {
	params := ShapeParameters{"side": 5}
	obj, err := NewShapeObject("square", "square", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ishape := obj.getIShape()
	square, ok := ishape.(*Square)
	if !ok {
		t.Fatalf("expected *Square, got %T", ishape)
	}
	if square.Side != 5 {
		t.Errorf("expected side 5, got %d", square.Side)
	}
}

func TestShapeObject_getIShape_Rectangle(t *testing.T) {
	params := ShapeParameters{"width": 3, "height": 7}
	obj, err := NewShapeObject("rectangle", "rectangle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ishape := obj.getIShape()
	rect, ok := ishape.(*Rectangle)
	if !ok {
		t.Fatalf("expected *Rectangle, got %T", ishape)
	}
	if rect.Width != 3 || rect.Height != 7 {
		t.Errorf("expected width 3 and height 7, got %d and %d", rect.Width, rect.Height)
	}
}

func TestShapeObject_getIShape_Triangle(t *testing.T) {
	params := ShapeParameters{"base": 4, "height": 6}
	obj, err := NewShapeObject("triangle", "triangle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ishape := obj.getIShape()
	tri, ok := ishape.(*Triangle)
	if !ok {
		t.Fatalf("expected *Triangle, got %T", ishape)
	}
	if tri.Base != 4 || tri.Height != 6 {
		t.Errorf("expected base 4 and height 6, got %d and %d", tri.Base, tri.Height)
	}
}

func TestShapeObject_getIShape_UnknownType(t *testing.T) {
	obj := &ShapeObject{
		ShapeType: "hexagon",
		Shape:     nil,
	}
	ishape := obj.getIShape()
	if ishape != nil {
		t.Errorf("expected nil for unknown shape type, got %T", ishape)
	}
}
func TestShapeObject_GetParameters_Circle(t *testing.T) {
	params := ShapeParameters{"radius": 10}
	obj, err := NewShapeObject("circle desc", "circle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := obj.GetParameters()
	if got["radius"] != 10 {
		t.Errorf("expected radius 10, got %d", got["radius"])
	}
}

func TestShapeObject_GetParameters_Square(t *testing.T) {
	params := ShapeParameters{"side": 4}
	obj, err := NewShapeObject("square desc", "square", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := obj.GetParameters()
	if got["side"] != 4 {
		t.Errorf("expected side 4, got %d", got["side"])
	}
}

func TestShapeObject_GetParameters_Rectangle(t *testing.T) {
	params := ShapeParameters{"width": 3, "height": 8}
	obj, err := NewShapeObject("rectangle desc", "rectangle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := obj.GetParameters()
	if got["width"] != 3 {
		t.Errorf("expected width 3, got %d", got["width"])
	}
	if got["height"] != 8 {
		t.Errorf("expected height 8, got %d", got["height"])
	}
}

func TestShapeObject_GetParameters_Triangle(t *testing.T) {
	params := ShapeParameters{"base": 5, "height": 12}
	obj, err := NewShapeObject("triangle desc", "triangle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := obj.GetParameters()
	if got["base"] != 5 {
		t.Errorf("expected base 5, got %d", got["base"])
	}
	if got["height"] != 12 {
		t.Errorf("expected height 12, got %d", got["height"])
	}
}

// Perimeter tests
func TestShapeObject_GetPerimeter_Circle(t *testing.T) {
	params := ShapeParameters{"radius": 3}
	obj, err := NewShapeObject("circle perimeter", "circle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 2 * 3.14 * 3 // 2πr
	got := obj.GetPerimeter()
	if got != expected {
		t.Errorf("expected perimeter %.2f, got %.2f", expected, got)
	}
}

func TestShapeObject_GetPerimeter_Square(t *testing.T) {
	params := ShapeParameters{"side": 4}
	obj, err := NewShapeObject("square perimeter", "square", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 16.0 // 4 * side
	got := obj.GetPerimeter()
	if got != expected {
		t.Errorf("expected perimeter %.2f, got %.2f", expected, got)
	}
}

func TestShapeObject_GetPerimeter_Rectangle(t *testing.T) {
	params := ShapeParameters{"width": 5, "height": 2}
	obj, err := NewShapeObject("rectangle perimeter", "rectangle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 2 * float64(5+2) // 2*(width+height)
	got := obj.GetPerimeter()
	if got != expected {
		t.Errorf("expected perimeter %.2f, got %.2f", expected, got)
	}
}

func TestShapeObject_GetPerimeter_Triangle(t *testing.T) {
	params := ShapeParameters{"base": 8, "height": 3}
	obj, err := NewShapeObject("triangle perimeter", "triangle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := float64(3 * 8) // 3 * base (equilateral)
	got := obj.GetPerimeter()
	if got != expected {
		t.Errorf("expected perimeter %.2f, got %.2f", expected, got)
	}
}

// Area tests
func TestShapeObject_GetArea_Circle(t *testing.T) {
	params := ShapeParameters{"radius": 3}
	obj, err := NewShapeObject("circle area", "circle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 3.14 * float64(3*3) // π * r^2
	got := obj.GetArea()
	if got != expected {
		t.Errorf("expected area %.2f, got %.2f", expected, got)
	}
}

func TestShapeObject_GetArea_Square(t *testing.T) {
	params := ShapeParameters{"side": 4}
	obj, err := NewShapeObject("square area", "square", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := float64(4 * 4) // side^2
	got := obj.GetArea()
	if got != expected {
		t.Errorf("expected area %.2f, got %.2f", expected, got)
	}
}

func TestShapeObject_GetArea_Rectangle(t *testing.T) {
	params := ShapeParameters{"width": 5, "height": 2}
	obj, err := NewShapeObject("rectangle area", "rectangle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := float64(5 * 2) // width * height
	got := obj.GetArea()
	if got != expected {
		t.Errorf("expected area %.2f, got %.2f", expected, got)
	}
}

func TestShapeObject_GetArea_Triangle(t *testing.T) {
	params := ShapeParameters{"base": 8, "height": 3}
	obj, err := NewShapeObject("triangle area", "triangle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 0.5 * float64(8*3) // 0.5 * base * height
	got := obj.GetArea()
	if got != expected {
		t.Errorf("expected area %.2f, got %.2f", expected, got)
	}
}

// Additional tests for edge cases and invalid usage

func TestShapeObject_GetArea_InvalidShapeType(t *testing.T) {
	obj := &ShapeObject{
		ShapeType: "hexagon",
		Shape:     nil,
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic or error when calling GetArea on invalid shape type")
		}
	}()
	_ = obj.GetArea()
}

func TestShapeObject_GetArea_NilShape(t *testing.T) {
	obj := &ShapeObject{
		ShapeType: "circle",
		Shape:     nil,
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic or error when calling GetArea with nil Shape")
		}
	}()
	_ = obj.GetArea()
}

func TestShapeObject_GetArea_ZeroParameters(t *testing.T) {
	params := ShapeParameters{"radius": 0}
	obj, err := NewShapeObject("zero circle", "circle", Pointer{0, 0}, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 0.0
	got := obj.GetArea()
	if got != expected {
		t.Errorf("expected area %.2f, got %.2f", expected, got)
	}
}
