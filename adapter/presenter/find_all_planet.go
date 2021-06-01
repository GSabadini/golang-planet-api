package presenter

import (
	"github.com/GSabadini/golang-planet-api/domain"
)

// FindAllPlanetOutput output data
type FindAllPlanetOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Climate   string `json:"climate"`
	Terrain   string `json:"terrain"`
	CreatedAt string `json:"created_at"`
}

func NewFindAllPlanetOutput() FindAllPlanetOutput {
	return FindAllPlanetOutput{}
}

func (f FindAllPlanetOutput) Output(planets []domain.Planet) []FindAllPlanetOutput {
	var output = make([]FindAllPlanetOutput, 0)

	for _, planet := range planets {
		output = append(output, FindAllPlanetOutput{
			ID:      planet.ID(),
			Name:    planet.Name(),
			Climate: planet.Climate(),
			Terrain: planet.Terrain(),
		})
	}
	return output
}
