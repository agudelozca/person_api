package services

/*
Implementa la interfaz PersonService con operaciones
específicas relacionadas con la entidad "Person".
Utiliza el repositorio PersonRepository para interactuar con la capa de almacenamiento.
*/

import (
	"github.com/agudelozca/person_api/models"
	"github.com/agudelozca/person_api/repositories"
)

type personService struct {
	PersonRepository repositories.PersonRepository
}

// NewPersonService crea una nueva instancia de PersonService.
func NewPersonService(personRepository repositories.PersonRepository) PersonService {
	return &personService{PersonRepository: personRepository}
}

func (s *personService) GetAllPeople() ([]models.Person, error) {
	return s.PersonRepository.GetAllPeople()
}

func (s *personService) GetPersonByID(id string) (models.Person, error) {
	return s.PersonRepository.GetPersonByID(id)
}

func (s *personService) CreatePerson(request CreatePersonRequest) (models.Person, error) {
	person := models.Person{
		Name: request.Name,
		Age:  request.Age,
	}

	return s.PersonRepository.CreatePerson(person)
}
