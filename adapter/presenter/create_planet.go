package presenter

import "github.com/GSabadini/golang-planet-api/domain"

type (
	// CreatePlanetOutput output data
	CreatePlanetOutput struct {
		ID        string                  `json:"id"`
		Name      string                  `json:"name"`
		Climate   string                  `json:"climate"`
		Terrain   string                  `json:"terrain"`
		Films     CreatePlanetFilmsOutput `json:"films"`
		CreatedAt string                  `json:"created_at"`
	}

	// CreatePlanetFilmsOutput output data
	CreatePlanetFilmsOutput struct {
		AppearedIn int `json:"appeared_in"`
	}
)

func NewCreatePlanetOutput() CreatePlanetOutput {
	return CreatePlanetOutput{}
}

func (f CreatePlanetOutput) Output(planet domain.Planet) CreatePlanetOutput {
	return CreatePlanetOutput{
		ID:      planet.ID(),
		Name:    planet.Name(),
		Climate: planet.Climate(),
		Films: CreatePlanetFilmsOutput{
			AppearedIn: planet.AppearedInFilms(),
		},
		Terrain: planet.Terrain(),
	}
}
