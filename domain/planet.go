package domain

import "context"

type (
	Planet struct {
		id      string
		name    string
		climate string
		ground  string
	}

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
)
