package repositories

import (
	"github.com/agudelozca/person_api/database"
	"github.com/agudelozca/person_api/models"
)

type personRepository struct{}

// NewPersonRepository crea una nueva instancia de PersonRepository.
func NewPersonRepository() PersonRepository {
	return &personRepository{}
}

func (r *personRepository) GetAllPeople() ([]models.Person, error) {
	return database.GetAllPeople()
}

func (r *personRepository) GetPersonByID(id string) (models.Person, error) {
	return database.GetPersonByID(id)
}

func (r *personRepository) CreatePerson(person models.Person) (models.Person, error) {
	return database.CreatePerson(person)
}
