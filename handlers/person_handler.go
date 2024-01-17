package handlers

/*
Maneja las solicitudes HTTP y delega la lógica de negocio
al servicio correspondiente.
Utiliza el servicio PersonService para realizar operaciones en la entidad "Person".
Serializa y deserializa datos JSON para las solicitudes y respuestas HTTP.*/

import (
	"encoding/json"
	"net/http"

	"github.com/agudelozca/person_api/services"
	"github.com/go-chi/chi/v5"
)

// PersonHandler maneja las solicitudes relacionadas con la entidad "Person".
type PersonHandler struct {
	PersonService services.PersonService
}

// NewPersonHandler crea una nueva instancia de PersonHandler.
func NewPersonHandler(personService services.PersonService) *PersonHandler {
	return &PersonHandler{PersonService: personService}
}

// GetAllPeople devuelve todos los registros en la base de datos.
func (h *PersonHandler) GetAllPeople(w http.ResponseWriter, r *http.Request) {
	people, err := h.PersonService.GetAllPeople()
	if err != nil {
		http.Error(w, "Error getting people", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// GetPersonByID devuelve un registro específico por ID.
func (h *PersonHandler) GetPersonByID(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "id")

	person, err := h.PersonService.GetPersonByID(personID)
	if err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

// CreatePerson agrega un nuevo registro a la base de datos.
func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var newPerson services.CreatePersonRequest
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdPerson, err := h.PersonService.CreatePerson(newPerson)
	if err != nil {
		http.Error(w, "Error creating person", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPerson)
}
