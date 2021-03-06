package presenter

import "github.com/GSabadini/golang-planet-api/domain"

// FindPlanetByIDOutput output data
type FindPlanetByIDOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Climate   string `json:"climate"`
	Terrain   string `json:"terrain"`
	CreatedAt string `json:"created_at"`
}

func NewFindPlanetByIDOutput() FindPlanetByIDOutput {
	return FindPlanetByIDOutput{}
}

func (f FindPlanetByIDOutput) Output(planet domain.Planet) FindPlanetByIDOutput {
	return FindPlanetByIDOutput{
		ID:      planet.ID(),
		Name:    planet.Name(),
		Climate: planet.Climate(),
		Terrain: planet.Terrain(),
	}
}
