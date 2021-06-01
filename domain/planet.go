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

	// PlanetFinderAll defines the search operation for a planet entity
	PlanetFinderAll interface {
		FindAll(context.Context) ([]Planet, error)
	}

	// PlanetFinderByName defines the search operation for a planet entity
	PlanetFinderByName interface {
		FindByName(context.Context, string) (Planet, error)
	}

	// PlanetFinderByID defines the search operation for a planet entity
	PlanetFinderByID interface {
		FindByID(context.Context, string) (Planet, error)
	}

	// PlanetDeleter defines the operation of removing a planet entity
	PlanetDeleter interface {
		Delete(context.Context, string) error
	}

	// Planet defines the planet entity
	Planet struct {
		id        string
		name      string
		climate   string
		terrain   string
		film     Film
		createdAt time.Time
	}

	// Film defines films property
	Film struct {
		appearedIn int
	}
)

// NewPlanet creates new Planet
func NewPlanet(id string, name string, climate string, terrain string, time time.Time) Planet {
	return Planet{
		id:        id,
		name:      name,
		climate:   climate,
		terrain:   terrain,
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

// Terrain returns the terrain property
func (p Planet) Terrain() string {
	return p.terrain
}

// CreatedAt returns the createdAt property
func (p Planet) CreatedAt() time.Time {
	return p.createdAt
}
