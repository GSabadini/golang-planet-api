package presenter

import "github.com/GSabadini/golang-planet-api/domain"

// CreatePlanetOutput output data
type CreatePlanetOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Climate   string `json:"climate"`
	Ground    string `json:"ground"`
	CreatedAt string `json:"created_at"`
}

func NewCreatePlanetOutput() FindPlanetByIDOutput {
	return FindPlanetByIDOutput{}
}

func (f CreatePlanetOutput) Output(planet domain.Planet) CreatePlanetOutput {
	return CreatePlanetOutput{
		ID:      planet.ID(),
		Name:    planet.Name(),
		Climate: planet.Climate(),
		Ground:  planet.Ground(),
	}
}
