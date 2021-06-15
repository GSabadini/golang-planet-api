package handler

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/GSabadini/golang-planet-api/domain"
	"github.com/GSabadini/golang-planet-api/usecase"

	"github.com/stretchr/testify/assert"
)

type stubCreatePlanetUseCase struct {
	result domain.Planet
	err    error
}

func (s stubCreatePlanetUseCase) Execute(_ context.Context, _ usecase.CreatePlanetInput) (domain.Planet, error) {
	return s.result, s.err
}

func TestCreatePlanetHandler_Handle(t *testing.T) {
	type fields struct {
		uc  usecase.CreatePlanetUseCase
		log *log.Logger
	}

	tests := []struct {
		name           string
		fields         fields
		rawPayload     []byte
		wantBody       string
		wantStatusCode int
	}{
		{
			name: "Create account successfully",
			fields: fields{
				uc: stubCreatePlanetUseCase{
					result: domain.NewPlanet(
						"fakeID",
						"Tatooine",
						"Arid",
						"Dessert",
						domain.NewFilms(0),
						time.Time{},
					),
					err: nil,
				},
				log: log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds),
			},
			rawPayload:     []byte(`{"name":"Tatooine","climate":"Arid","terrain":"Dessert"}`),
			wantBody:       `{"id":"fakeID","name":"Tatooine","climate":"Arid","terrain":"Dessert","films":{"appeared_in":0},"created_at":""}`,
			wantStatusCode: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPost,
				"/planet",
				bytes.NewReader(tt.rawPayload),
			)
			if err != nil {
				t.Fatal(err)
			}

			var (
				rr      = httptest.NewRecorder()
				handler = NewCreatePlanetHandler(tt.fields.uc, tt.fields.log)
			)

			handler.Handle(rr, req)

			var (
				gotStatusCode = rr.Code
				gotPayload    = rr.Body.String()
			)

			if assert.Equal(t, tt.wantStatusCode, gotStatusCode) {
				assert.JSONEq(t, tt.wantBody, gotPayload)
			}
		})
	}
}
