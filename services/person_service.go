package services

import "github.com/agudelozca/person_api/models"

// PersonService define las operaciones relacionadas con la entidad "Person".
type PersonService interface {
	GetAllPeople() ([]models.Person, error)
	GetPersonByID(id string) (models.Person, error)
	CreatePerson(request CreatePersonRequest) (models.Person, error)
}

// CreatePersonRequest define la estructura para la creación de una persona.
type CreatePersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
