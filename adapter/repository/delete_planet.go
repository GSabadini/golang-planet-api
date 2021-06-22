package handler

import (
	"context"

	"github.com/GSabadini/golang-planet-api/domain"
	"github.com/GSabadini/golang-planet-api/infrastructure/database"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	deleteResult = 0
)

type (
	deletePlanetRepository struct {
		handler    *database.MongoHandler
		collection string
	}
)

// NewDeletePlanetRepository creates new deletePlanetRepository with its dependencies
func NewDeletePlanetRepository(handler *database.MongoHandler, collection string) domain.PlanetDeleter {
	return deletePlanetRepository{handler: handler, collection: collection}
}

// Delete performs DeleteOne into the database
func (d deletePlanetRepository) Delete(ctx context.Context, ID string) error {
	var query = bson.D{{Key: "id", Value: ID}}

	result, err := d.handler.DB().Collection(d.collection).DeleteOne(ctx, query)
	if err != nil {
		return err
	}

	if result.DeletedCount == deleteResult {
		return domain.ErrPlanetNotFound
	}

	return nil
}
