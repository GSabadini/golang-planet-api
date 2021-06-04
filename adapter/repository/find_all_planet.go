package handler

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
	"github.com/GSabadini/golang-planet-api/infrastructure/database"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	// Bson data
	findAllPlanetBSON struct {
		ID        string                 `bson:"id"`
		Name      string                 `bson:"name"`
		Climate   string                 `bson:"climate"`
		Terrain   string                 `bson:"terrain"`
		Films     findAllPlanetFilmsBSON `bson:"Films"`
		CreatedAt time.Time              `bson:"created_at"`
	}

	// Bson data
	findAllPlanetFilmsBSON struct {
		AppearedIn int `bson:"appeared_in"`
	}

	findAllPlanetRepository struct {
		handler    *database.MongoHandler
		collection string
	}
)

// NewFindAllPlanetRepository creates new findAllPlanetRepository with its dependencies
func NewFindAllPlanetRepository(handler *database.MongoHandler, collection string) domain.PlanetFinderAll {
	return findAllPlanetRepository{handler: handler, collection: collection}
}

// FindAll performs Find into the database
func (f findAllPlanetRepository) FindAll(ctx context.Context) ([]domain.Planet, error) {
	var (
		planetsBSON = make([]findAllPlanetBSON, 0)
		planets     = make([]domain.Planet, 0)
	)

	cur, err := f.handler.DB().Collection(f.collection).Find(ctx, bson.D{})
	if err != nil {
		return []domain.Planet{}, errors.Wrap(err, domain.ErrFindAllPlanet.Error())
	}

	defer cur.Close(ctx)
	if err = cur.All(ctx, &planetsBSON); err != nil {
		return []domain.Planet{}, errors.Wrap(err, domain.ErrFindAllPlanet.Error())
	}

	if err := cur.Err(); err != nil {
		return []domain.Planet{}, errors.Wrap(err, domain.ErrFindAllPlanet.Error())
	}

	for _, planetBSON := range planetsBSON {
		planets = append(planets, domain.NewPlanet(
			planetBSON.ID,
			planetBSON.Name,
			planetBSON.Climate,
			planetBSON.Terrain,
			domain.NewFilms(planetBSON.Films.AppearedIn),
			planetBSON.CreatedAt,
		),
		)
	}

	return planets, nil
}
