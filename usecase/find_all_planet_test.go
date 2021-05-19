package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
)

type stubPlanetFinderRepository struct {
	domain.PlanetFinder
	result []domain.Planet
	err    error
}

func (s stubPlanetFinderRepository) FindAll(_ context.Context) ([]domain.Planet, error) {
	return s.result, s.err
}

type stubPlanetFinderPresenter struct{}

func (s stubPlanetFinderPresenter) Output(planets []domain.Planet) []FindAllPlanetOutput {
	var output = make([]FindAllPlanetOutput, 0)
	for _, planet := range planets {
		output = append(output, FindAllPlanetOutput{
			ID:      planet.ID(),
			Name:    planet.Name(),
			Climate: planet.Climate(),
			Ground:  planet.Ground(),
		})
	}
	return output
}
func Test_findAllPlanetInteractor_Execute(t *testing.T) {
	type fields struct {
		repository domain.PlanetFinder
		presenter  FindAllPlanetPresenter
		ctxTimeout time.Duration
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []FindAllPlanetOutput
		wantErr bool
	}{
		{
			name: "Should return a list of all the planets",
			fields: fields{
				repository: stubPlanetFinderRepository{
					result: []domain.Planet{
						domain.NewPlanet(
							"fakeID",
							"fakeName",
							"fakeClimate",
							"fakeGround",
							time.Time{},
						),
						domain.NewPlanet(
							"fakeID2",
							"fakeName2",
							"fakeClimate2",
							"fakeGround2",
							time.Time{},
						),
					},
					err: nil,
				},
				presenter:  stubPlanetFinderPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
			},
			want: []FindAllPlanetOutput{
				{
					ID:      "fakeID",
					Name:    "fakeName",
					Climate: "fakeClimate",
					Ground:  "fakeGround",
				},
				{
					ID:      "fakeID2",
					Name:    "fakeName2",
					Climate: "fakeClimate2",
					Ground:  "fakeGround2",
				},
			},
			wantErr: false,
		},
		{
			name: "Should fail to return a list of all planets",
			fields: fields{
				repository: stubPlanetFinderRepository{
					result: []domain.Planet{},
					err:    errors.New("failed to create the planet"),
				},
				presenter:  stubPlanetFinderPresenter{},
				ctxTimeout: 0,
			},
			args: args{
				ctx: context.TODO(),
			},
			want:    []FindAllPlanetOutput{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFindAllPlanetInteractor(
				tt.fields.repository,
				tt.fields.presenter,
				tt.fields.ctxTimeout,
			)

			got, err := f.Execute(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}
