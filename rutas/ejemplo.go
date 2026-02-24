package rutas

import (
	"html/template"
	"net/http"
	"proyecto_3/utilidades"
	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) { 
	// En algunos codigos se le llama w (response) y r (request), pero es lo mismo
	
	template := template.Must(template.ParseFiles("templates/ejemplo/home.html", utilidades.Frontend))
	template.Execute(response, nil)

	// La función Must se utiliza para manejar errores de forma más sencilla. Si ocurre un error al analizar los archivos de plantilla, 
	// Must hará que el programa se detenga y muestre el error en lugar de devolverlo como un valor. 
	// Esto es útil para asegurarse de que las plantillas se carguen correctamente antes de continuar con la ejecución del programa.

	/*	Codigo sin usar template.Must
		template, err := template.ParseFiles("templates/ejemplo/home.html", "templates/layout/base.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(response, nil)
		}
	*/
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/nosotros.html", utilidades.Frontend))
	template.Execute(response, nil)
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/parametros.html", utilidades.Frontend))
	vars := mux.Vars(request)
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
	}
		template.Execute(response, data)
}

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/parametros_querystring.html", utilidades.Frontend))
	data := map[string]string{
		"id": 	request.URL.Query().Get("id"),
		"slug": request.URL.Query().Get("slug"),
	}
		template.Execute(response, data)
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

func Estructuras(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/ejemplo/estructuras.html", utilidades.Frontend))
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

func Pagina404(response http.ResponseWriter, request *http.Request) {

	template := template.Must(template.ParseFiles("templates/ejemplo/pagina404.html", utilidades.Frontend))
	template.Execute(response, nil)

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