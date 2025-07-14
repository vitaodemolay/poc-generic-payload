package containers

import (
	domain "github.com/vitaodemolay/poc-generic-payload/internal/domain/shapes"
	"github.com/vitaodemolay/poc-generic-payload/internal/infrastructure/nosql/shapes"
)

type Infrastructure struct {
	ShapeRepository domain.ShapeRepository
}

const (
	DefaultMongoConnectionString = "mongodb://admin:PassW0rd!@localhost:27017"
	DefaultMongoDatabase         = "shapesdb"
	DefaultMongoCollection       = "shapes"
)

func NewInfrastructure() (*Infrastructure, error) {
	repo, err := shapes.NewShapeRepository(
		DefaultMongoConnectionString,
		DefaultMongoDatabase,
		DefaultMongoCollection,
	)
	if err != nil {
		return nil, err
	}

	return &Infrastructure{
		ShapeRepository: repo,
	}, nil
}
