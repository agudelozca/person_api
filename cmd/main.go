package main

/*
Inicializa la base de datos al arrancar la aplicación.
Inyecta las dependencias necesarias para el controlador,
el servicio y el repositorio.
Configura las rutas de la API utilizando la biblioteca chi.
*/

import (
	"fmt"
	"net/http"

	"github.com/agudelozca/person_api/database"
	"github.com/agudelozca/person_api/handlers"
	"github.com/agudelozca/person_api/repositories"
	"github.com/agudelozca/person_api/services"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	// Inicializar la base de datos al iniciar la aplicación.
	err := database.Init("./database.json")
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	// Inyectar dependencias
	personRepository := repositories.NewPersonRepository()
	personService := services.NewPersonService(personRepository)
	personHandler := handlers.NewPersonHandler(personService)

	// Configurar rutas
	router.Get("/people", personHandler.GetAllPeople)
	router.Get("/people/{id}", personHandler.GetPersonByID)
	router.Post("/people", personHandler.CreatePerson)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}
