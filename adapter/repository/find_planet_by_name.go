package handler

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
	"github.com/GSabadini/golang-planet-api/infrastructure/database"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// Bson data
	findPlanetByNameBSON struct {
		ID        string                    `bson:"id"`
		Name      string                    `bson:"name"`
		Climate   string                    `bson:"climate"`
		Terrain   string                    `bson:"terrain"`
		Films     findPlanetByNameFilmsBSON `bson:"Films"`
		CreatedAt time.Time                 `bson:"created_at"`
	}

	// Bson data
	findPlanetByNameFilmsBSON struct {
		AppearedIn int `bson:"appeared_in"`
	}

	findPlanetByNameRepository struct {
		handler    *database.MongoHandler
		collection string
	}
)

// NewFindPlanetByNameRepository creates new findPlanetByNameRepository with its dependencies
func NewFindPlanetByNameRepository(handler *database.MongoHandler, collection string) domain.PlanetFinderByName {
	return findPlanetByNameRepository{handler: handler, collection: collection}
}

// FindByName performs FindOne into the database
func (f findPlanetByNameRepository) FindByName(ctx context.Context, name string) (domain.Planet, error) {
	var (
		planetBSON = findPlanetByNameBSON{}
		query      = bson.M{"name": name}
	)

	var err = f.handler.DB().Collection(f.collection).FindOne(ctx, query).Decode(&planetBSON)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return domain.Planet{}, domain.ErrPlanetNotFound
		default:
			return domain.Planet{}, errors.Wrap(err, domain.ErrFindPlanetByName.Error())
		}
	}

	return domain.NewPlanet(
		planetBSON.ID,
		planetBSON.Name,
		planetBSON.Climate,
		planetBSON.Terrain,
		domain.NewFilms(planetBSON.Films.AppearedIn),
		planetBSON.CreatedAt,
	), nil
}
