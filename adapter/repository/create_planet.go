package handler

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
	"github.com/GSabadini/golang-planet-api/infrastructure/database"

	"github.com/pkg/errors"
)

type (
	// Bson data
	createPlanetBSON struct {
		ID        string                `bson:"id"`
		Name      string                `bson:"name"`
		Climate   string                `bson:"climate"`
		Terrain   string                `bson:"terrain"`
		films     createPlanetFilmsBSON `bson:"Films"`
		CreatedAt time.Time             `bson:"created_at"`
	}

	// Bson data
	createPlanetFilmsBSON struct {
		AppearedIn int `bson:"appeared_in"`
	}

	createPlanetRepository struct {
		handler    *database.MongoHandler
		collection string
	}
)

// NewCreatePlanetRepository creates new createPlanetRepository with its dependencies
func NewCreatePlanetRepository(handler *database.MongoHandler, collection string) domain.PlanetCreator {
	return createPlanetRepository{handler: handler, collection: collection}
}

// Create performs InsertOne into the database
func (c createPlanetRepository) Create(ctx context.Context, planet domain.Planet) (domain.Planet, error) {
	var bson = createPlanetBSON{
		ID:      planet.ID(),
		Name:    planet.Name(),
		Climate: planet.Climate(),
		Terrain: planet.Terrain(),
		films: createPlanetFilmsBSON{
			AppearedIn: planet.AppearedInFilms(),
		},
		CreatedAt: planet.CreatedAt(),
	}

	if _, err := c.handler.DB().Collection(c.collection).InsertOne(ctx, bson); err != nil {
		return domain.Planet{}, errors.Wrap(err, domain.ErrCreatePlanet.Error())
	}

	return planet, nil
}
