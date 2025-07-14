package contracts

import "github.com/vitaodemolay/poc-generic-payload/internal/domain/shapes"

type CreateShape struct {
	Description string                 `json:"description" validate:"required"`
	ShapeType   string                 `json:"shapeType" validate:"required,oneof=circle square rectangle triangle"`
	StartPoint  shapes.Pointer         `json:"startPoint" validate:"required"`
	Parameters  shapes.ShapeParameters `json:"parameters" validate:"required"`
}

type Shape struct {
	ID          string                 `json:"id"`
	Description string                 `json:"description"`
	ShapeType   string                 `json:"shapeType"`
	ActualPoint shapes.Pointer         `json:"actualPoint"`
	Parameters  shapes.ShapeParameters `json:"parameters"`
	Area        float64                `json:"area"`
	Perimeter   float64                `json:"perimeter"`
}

type ChangeShapePosition struct {
	ID       string         `json:"id" validate:"required"`
	NewPoint shapes.Pointer `json:"newPoint" validate:"required"`
}
