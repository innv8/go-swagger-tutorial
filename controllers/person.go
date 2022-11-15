package controllers

import (
	"encoding/json"
	"fmt"
	"go-swagger-tutorial/entities"
	"go-swagger-tutorial/middleware"
	"go-swagger-tutorial/models"
	"log"
	"net/http"
)

// ErrorResponse model info
// @Description The generic error messages for all endpoints
type ErrorResponse struct {
	Message string `json:"message" example:"Your account was rejected"`
}

// AddPersonPayload is the request to add a new person
// @Description Payload used to create new person
type AddPersonPayload struct {
	Name string `json:"name" example:"John Doe"`
	Age  uint   `json:"age" example:"34"`
}

func (p *AddPersonPayload) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if p.Age <= 0 {
		return fmt.Errorf("age is required")
	}
	return nil
}

// AddPersonResponse
// @Description the successful response of the created user
type AddPersonResponse struct {
	Id int `json:"id" example:"3"`
}

// AddPerson godoc
// @Summary		Adds a new person
// @Description	Make a POST request to /persons to create a new person
// @Tags 		person
// @Accept 		json
// @Produce 	json
// @Param		data	body	AddPersonPayload true	"The person input"
// @Success 	200		{object} AddPersonResponse
// @Failure		400		{object} ErrorResponse
// @Router		/persons	[post]
func (b *Base) AddPerson(w http.ResponseWriter, r *http.Request) {
	var data AddPersonPayload
	var response AddPersonResponse
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("unable to decode data because %v", err)
		middleware.JSONResponse(w, 400, ErrorResponse{Message: "invalid data sent"})
		return
	}

	err = data.Validate()
	if err != nil {
		log.Printf("invalid data : %v", err)
		middleware.JSONResponse(w, 400, ErrorResponse{Message: err.Error()})
		return
	}
	response.Id = models.AddPerson(data.Name, data.Age)

	middleware.JSONResponse(w, 200, response)
}

type GetPersonsResponse struct {
	Data []entities.Person `json:"data"`
}

func (b *Base) GetPersons(w http.ResponseWriter, _ *http.Request) {
	var response GetPersonsResponse

	response.Data = models.GetPersons()
	middleware.JSONResponse(w, 200, response)
}
