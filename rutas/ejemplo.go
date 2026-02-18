package rutas

import (
	//"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) { // En algunos codigos se le llama w (response) y r (request), pero es lo mismo
	template, err := template.ParseFiles("templates/ejemplo/home.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/nosotros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/parametros.html")
	vars := mux.Vars(request)
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/ejemplo/parametros_querystring.html")
	data := map[string]string{
		"id": 	request.URL.Query().Get("id"),
		"slug": request.URL.Query().Get("slug"),
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}
type Habilidad struct {
	Nombre string
}

type Datos struct {
	Nombre string
	Edad int
	Perfil int
	Habilidades []Habilidad
}

func Estructuras(response http.ResponseWriter, request *http.Request) { // En algunos codigos se le llama w (response) y r (request), pero es lo mismo
	template, _ := template.ParseFiles("template/ejemplo/estructuras.html")
	habilidad1 := Habilidad{"Inteligencia"}
	habilidad2 := Habilidad{"Deportes"}
	habilidad3 := Habilidad{"Cantar"}
	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}

	template.Execute(response, Datos{"Norman", 30, 2, habilidades})
	/* 
		template, err := template.ParseFiles("templates/ejemplo/estructuras.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, Datos{"Norman", 30, 2})
		} 
	*/
}

/* 
func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "hola mundo desde golang")
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Test en la terminal con fresh")
	fmt.Fprintln(response, "Sobre Nosotros")
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
 */