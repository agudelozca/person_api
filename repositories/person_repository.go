package repositories

import "github.com/agudelozca/person_api/models"

// PersonRepository define las operaciones de almacenamiento y recuperación de datos relacionadas con la entidad "Person".
type PersonRepository interface {
	GetAllPeople() ([]models.Person, error)
	GetPersonByID(id string) (models.Person, error)
	CreatePerson(person models.Person) (models.Person, error)
}
