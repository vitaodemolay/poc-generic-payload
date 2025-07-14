package containers

import (
	"github.com/vitaodemolay/poc-generic-payload/internal/application/shapes"
	domain "github.com/vitaodemolay/poc-generic-payload/internal/domain/shapes"
)

type Application struct {
	ShapeService shapes.ShapeService
}

func NewApplicationWithInjection(infraContainer *Infrastructure) (*Application, error) {
	repo := infraContainer.ShapeRepository
	shapeService := shapes.NewShapeService(repo)

	return &Application{
		ShapeService: shapeService,
	}, nil
}

func NewApplication() (*Application, error) {
	repo := NewFakeShapeRepository()
	shapeService := shapes.NewShapeService(repo)

	return &Application{
		ShapeService: shapeService,
	}, nil
}

// Fake Respository for testing purposes
type FakeShapeRepository struct {
	shapes map[string]*domain.ShapeObject
}

func NewFakeShapeRepository() *FakeShapeRepository {
	return &FakeShapeRepository{
		shapes: make(map[string]*domain.ShapeObject),
	}
}

func (r *FakeShapeRepository) Save(shape *domain.ShapeObject) error {
	r.shapes[shape.ID] = shape
	return nil
}

func (r *FakeShapeRepository) GetByID(id string) (*domain.ShapeObject, error) {
	shape, exists := r.shapes[id]
	if !exists {
		return nil, nil // or return an error if preferred
	}
	return shape, nil
}
