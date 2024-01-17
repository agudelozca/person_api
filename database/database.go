package database

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"

	"github.com/agudelozca/person_api/models"
)

// Database representa la "base de datos" almacenada en un archivo JSON.
var Database = struct {
	sync.RWMutex
	People  map[string]models.Person
	Counter int
}{People: make(map[string]models.Person)}

// Init carga los datos desde el archivo JSON en la base de datos.
func Init(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Database.People)
	if err != nil {
		return err
	}

	return nil
}

// Save guarda los datos de la base de datos en el archivo JSON.
func Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(Database.People)
	if err != nil {
		return err
	}

	return nil
}

// GetAllPeople devuelve todos los registros en la base de datos.
func GetAllPeople() ([]models.Person, error) {
	Database.RLock()
	defer Database.RUnlock()

	peopleList := make([]models.Person, 0, len(Database.People))
	for _, person := range Database.People {
		peopleList = append(peopleList, person)
	}

	return peopleList, nil
}

// GetPersonByID devuelve un registro específico por ID.
func GetPersonByID(id string) (models.Person, error) {
	Database.RLock()
	defer Database.RUnlock()

	person, exists := Database.People[id]
	if !exists {
		return models.Person{}, nil // Devuelve una estructura de persona vacía si no se encuentra.
	}

	return person, nil
}

// CreatePerson agrega un nuevo registro a la base de datos.
func CreatePerson(person models.Person) (models.Person, error) {
	Database.Lock()
	defer Database.Unlock()

	// Generar un ID único incremental
	Database.Counter++
	person.ID = strconv.Itoa(Database.Counter)

	Database.People[person.ID] = person

	// Guardar la base de datos actualizada en el archivo JSON.
	err := Save("database.json")
	if err != nil {
		return models.Person{}, err
	}

	return person, nil
}
