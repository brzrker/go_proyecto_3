package rutas

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "hola mundo desde golang")
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Test en la terminal con fresh")
	fmt.Fprintln(response, "Sobre Nosotros con ñanduuu")
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ID = "+vars["id"]+" | Slug = "+vars["slug"])
}

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, request.URL) // Imprime la URL completa
	fmt.Fprintln(response, request.URL.RawQuery) // Imprime solo la parte de la query string
	fmt.Fprintln(response, request.URL.Query()) // Imprime un mapa con los parámetros de la query string, por ejemplo: map[name:[John] age:[30]]
	fmt.Fprintln(response, request.URL.Query().Get("id")) // Imprime el valor del parámetro "id", por ejemplo: 123
	fmt.Fprintln(response, request.URL.Query().Get("slug")) // Imprime el valor del parámetro "slug"
	id := request.URL.Query().Get("id")
	slug := request.URL.Query().Get("slug")
	fmt.Fprintln(response, "-------------------------------------------")
	fmt.Fprintf(response, "id = %s | slug = %s", id, slug)
}