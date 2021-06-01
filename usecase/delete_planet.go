package usecase

import (
	"context"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type (
	// DeletePlanetUseCase input port
	DeletePlanetUseCase interface {
		Execute(context.Context, DeletePlanetInput) error
	}

	// DeletePlanetInput input data
	DeletePlanetInput struct {
		ID string
	}

	deletePlanetInteractor struct {
		repository domain.PlanetDeleter
		ctxTimeout time.Duration
	}
)

// NewDeletePlanetInteractor creates new deletePlanetInteractor with its dependencies
func NewDeletePlanetInteractor(
	repository domain.PlanetDeleter,
	ctxTimeout time.Duration,
) DeletePlanetUseCase {
	return deletePlanetInteractor{
		repository: repository,
		ctxTimeout: ctxTimeout,
	}
}

// Execute orchestrates the use case
func (c deletePlanetInteractor) Execute(ctx context.Context, input DeletePlanetInput) error {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout*time.Second)
	defer cancel()

	err := c.repository.Delete(ctx, input.ID)
	if err != nil {
		return err
	}

	return nil
}
