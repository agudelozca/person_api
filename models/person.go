package models

// Person representa la estructura de los datos almacenados en la base de datos.
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
