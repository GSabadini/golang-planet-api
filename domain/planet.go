package domain

import (
	"context"
	"time"
)

type (
	// PlanetCreator defines the operation of creating a account entity
	PlanetCreator interface {
		Create(context.Context, Planet) (Planet, error)
	}

	// PlanetFinder defines the search operation for a account entity
	PlanetFinder interface {
		FindAll(context.Context, string) ([]Planet, error)
		FindByID(context.Context, string) (Planet, error)
		FindByName(context.Context, string) (Planet, error)
	}

	// PlanetDeleter defines the search operation for a account entity
	PlanetDeleter interface {
		Delete(context.Context, int64) error
	}

	Planet struct {
		id        string
		name      string
		climate   string
		ground    string
		createdAt time.Time
	}
)

func NewPlanet(id string, name string, climate string, ground string, time time.Time) Planet {
	return Planet{
		id:        id,
		name:      name,
		climate:   climate,
		ground:    ground,
		createdAt: time,
	}
}

func (p Planet) ID() string {
	return p.id
}

func (p Planet) Name() string {
	return p.name
}

func (p Planet) Climate() string {
	return p.climate
}

func (p Planet) Ground() string {
	return p.ground
}

func (p Planet) CreatedAt() time.Time {
	return p.createdAt
}
