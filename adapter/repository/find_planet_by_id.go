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
	findPlanetByIDBSON struct {
		ID        string                  `bson:"id"`
		Name      string                  `bson:"name"`
		Climate   string                  `bson:"climate"`
		Terrain   string                  `bson:"terrain"`
		Films     findPlanetByIDFilmsBSON `bson:"Films"`
		CreatedAt time.Time               `bson:"created_at"`
	}

	// Bson data
	findPlanetByIDFilmsBSON struct {
		AppearedIn int `bson:"appeared_in"`
	}

	findPlanetByIDRepository struct {
		handler    *database.MongoHandler
		collection string
	}
)

// NewFindPlanetByIDRepository creates new findPlanetByIDRepository with its dependencies
func NewFindPlanetByIDRepository(handler *database.MongoHandler, collection string) domain.PlanetFinderByID {
	return findPlanetByIDRepository{handler: handler, collection: collection}
}

// FindByID performs FindOne into the database
func (f findPlanetByIDRepository) FindByID(ctx context.Context, ID string) (domain.Planet, error) {
	var (
		planetBSON = findPlanetByIDBSON{}
		query      = bson.M{"id": ID}
	)

	var err = f.handler.DB().Collection(f.collection).FindOne(ctx, query).Decode(&planetBSON)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return domain.Planet{}, domain.ErrPlanetNotFound
		default:
			return domain.Planet{}, errors.Wrap(err, domain.ErrFindPlanetByID.Error())
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
