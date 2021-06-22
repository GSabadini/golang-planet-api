package handler

import (
	"encoding/json"
	"github.com/GSabadini/golang-planet-api/adapter/presenter"
	"log"
	"net/http"

	"github.com/GSabadini/golang-planet-api/adapter/api/response"
	"github.com/GSabadini/golang-planet-api/usecase"

	"github.com/go-ozzo/ozzo-validation"
)

// CreatePlanetHandler defines the dependencies of the HTTP handler for the use case
type CreatePlanetHandler struct {
	uc  usecase.CreatePlanetUseCase
	log *log.Logger
}

// NewCreatePlanetHandler creates new CreatePlanetHandler with its dependencies
func NewCreatePlanetHandler(uc usecase.CreatePlanetUseCase, log *log.Logger) CreatePlanetHandler {
	return CreatePlanetHandler{
		uc:  uc,
		log: log,
	}
}

// Handle handles http request
func (c CreatePlanetHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreatePlanetInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.log.Println("failed to marshal message:", err)
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if err := c.Validate(input); err != nil {
		c.log.Println("invalid input:", err)
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := c.uc.Execute(r.Context(), input)
	if err != nil {
		c.log.Println("failed to creating account:", err)
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
		//switch {
		//case errors.Is(err, domain.ErrPlanetAlreadyExists):
		//	response.NewError(err, http.StatusUnprocessableEntity).Send(w)
		//	return
		//default:
		//	response.NewError(err, http.StatusInternalServerError).Send(w)
		//	return
		//}

	}

	c.log.Println("success to creating account")
	response.NewSuccess(presenter.NewCreatePlanetOutput().Output(output), http.StatusCreated).Send(w)
}

func (c CreatePlanetHandler) Validate(i usecase.CreatePlanetInput) error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&i.Climate, validation.Required, validation.Length(1, 100)),
		validation.Field(&i.Terrain, validation.Required, validation.Length(1, 100)),
	)
}
