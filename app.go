package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Creamos una estructura para usuarios
type Users struct {
	Name string
	Skills string
	Age int
}

type Year struct {
	Year int
}

func Index(rw http.ResponseWriter, r *http.Request) {

	// Retornar el template a la variable template, el _ sirve para no retornar el error
	//template, _ := template.ParseFiles("templates/index.html")
	template, err := template.ParseFiles("templates/index.html")

	// Generamos un usuario usando la estructura Users
	user := Users{"Ivan Cuenca", "Desarrollador Go", 42}

	// Si no genera error ejecuta el template, y le enviamos valores (user)
	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}
}

func About(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/about.html")
	year := Year{2023}
	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, year)
	}
}

func main() {

	// Aqui se agregan las rutas y los metodos a llamar
	http.HandleFunc("/", Index)
	http.HandleFunc("/about", About)
	
	// Crea el servidor
	iniciaServer()
}

func iniciaServer() {
	// Funcion para inicializar el servidor
	fmt.Println("********El servidor esta corriendo en puerto 3000********")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}