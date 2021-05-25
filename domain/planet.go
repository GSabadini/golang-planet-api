package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrPlanetNotFound = errors.New("planet not found")
)

type (
	// PlanetCreator defines the operation of creating a planet entity
	PlanetCreator interface {
		Create(context.Context, Planet) (Planet, error)
	}

	// PlanetFinder defines the search operation for a planet entity
	PlanetFinder interface {
		FindAll(context.Context) ([]Planet, error)
		FindByID(context.Context, string) (Planet, error)
		FindByName(context.Context, string) (Planet, error)
	}

	// PlanetDeleter defines the operation of removing a planet entity
	PlanetDeleter interface {
		Delete(context.Context, int64) error
	}

	// Planet defines the planet entity
	Planet struct {
		id        string
		name      string
		climate   string
		ground    string
		createdAt time.Time
	}
)

// NewPlanet creates new Planet
func NewPlanet(id string, name string, climate string, ground string, time time.Time) Planet {
	return Planet{
		id:        id,
		name:      name,
		climate:   climate,
		ground:    ground,
		createdAt: time,
	}
}

// ID returns the id property
func (p Planet) ID() string {
	return p.id
}

// Name returns the name property
func (p Planet) Name() string {
	return p.name
}

// Climate returns the climate property
func (p Planet) Climate() string {
	return p.climate
}

// Ground returns the ground property
func (p Planet) Ground() string {
	return p.ground
}

// CreatedAt returns the createdAt property
func (p Planet) CreatedAt() time.Time {
	return p.createdAt
}
