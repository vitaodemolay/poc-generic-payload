package shapes

import (
	"context"

	"github.com/vitaodemolay/poc-generic-payload/internal/application/shapes/contracts"
	"github.com/vitaodemolay/poc-generic-payload/internal/domain/shapes"
)

type ShapeService interface {
	CreateShape(ctx context.Context, dto contracts.CreateShape) (string, error)
	GetShape(ctx context.Context, id string) (contracts.Shape, error)
	ChangeShapePosition(ctx context.Context, dto contracts.ChangeShapePosition) error
}

type shapeService struct {
	repository shapes.ShapeRepository
}

func NewShapeService(repository shapes.ShapeRepository) ShapeService {
	return &shapeService{repository: repository}
}

// Implement the methods of ShapeService interface
func (s *shapeService) CreateShape(ctx context.Context, dto contracts.CreateShape) (string, error) {
	shape, err := shapes.NewShapeObject(dto.Description, dto.ShapeType, dto.StartPoint, dto.Parameters)
	if err != nil {
		return "", err
	}

	err = s.repository.Save(shape)
	if err != nil {
		return "", err
	}

	return shape.ID, nil
}

func (s *shapeService) GetShape(ctx context.Context, id string) (contracts.Shape, error) {
	shape, err := s.repository.GetByID(id)
	if err != nil {
		return contracts.Shape{}, err
	}

	return contracts.Shape{
		ID:          shape.ID,
		Description: shape.Description,
		ShapeType:   shape.ShapeType,
		ActualPoint: shape.ActualPoint,
		Parameters:  shape.GetParameters(),
		Area:        shape.GetArea(),
		Perimeter:   shape.GetPerimeter(),
	}, nil
}

func (s *shapeService) ChangeShapePosition(ctx context.Context, dto contracts.ChangeShapePosition) error {
	// Implementation logic for changing the position of a shape
	shape, err := s.repository.GetByID(dto.ID)
	if err != nil {
		return err
	}

	// Update the shape's position
	shape.ActualPoint = dto.NewPoint
	err = s.repository.Save(shape)
	if err != nil {
		return err
	}

	return nil
}
