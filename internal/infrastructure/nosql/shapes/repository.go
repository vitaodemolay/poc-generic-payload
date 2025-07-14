package shapes

import (
	"context"

	"github.com/vitaodemolay/poc-generic-payload/internal/domain/shapes"
	internalerrors "github.com/vitaodemolay/poc-generic-payload/pkg/internal-errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShapeRepository struct {
	mongoClient *mongo.Client
	database    string
	collection  string
}

func NewShapeRepository(connectionString, database, collection string) (*ShapeRepository, error) {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &ShapeRepository{
		mongoClient: client,
		database:    database,
		collection:  collection,
	}, nil
}

func (r *ShapeRepository) Save(shape *shapes.ShapeObject) error {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)
	filter := bson.M{"_id": shape.ID}
	options := options.Replace().SetUpsert(true)

	_, err := collection.ReplaceOne(
		context.Background(),
		filter,
		shape,
		options,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShapeRepository) GetByID(id string) (*shapes.ShapeObject, error) {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)
	filter := bson.M{"_id": id}

	var shape shapes.ShapeObject
	err := collection.FindOne(context.Background(), filter).Decode(&shape)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, internalerrors.ErrNotFound // Shape not found
		}
		return nil, err // Other error
	}

	return &shape, nil
}
