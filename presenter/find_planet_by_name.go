package presenter

import (
	"github.com/GSabadini/golang-planet-api/domain"
)

// FindPlanetByNameOutput output data
type FindPlanetByNameOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Climate   string `json:"climate"`
	Ground    string `json:"ground"`
	CreatedAt string `json:"created_at"`
}

func NewFindPlanetByNameOutput() FindPlanetByNameOutput {
	return FindPlanetByNameOutput{}
}

func (f FindPlanetByNameOutput) Output(planet domain.Planet) FindPlanetByNameOutput {
	return FindPlanetByNameOutput{
		ID:      planet.ID(),
		Name:    planet.Name(),
		Climate: planet.Climate(),
		Ground:  planet.Ground(),
	}
}
