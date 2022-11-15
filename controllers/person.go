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

type ErrorResponse struct {
	Message string `json:"message"`
}

type AddPersonPayload struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
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

type AddPersonResponse struct {
	Id int `json:"id"`
}

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
