package rutas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"proyecto_3/modelos"
	"proyecto_3/utilidades"
	"github.com/gorilla/mux"
)

var Token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MzYsImlhdCI6MTc3MjEzNTE0NSwiZXhwIjoxNzc0NzI3MTQ1fQ.iFwr87pDFHpOq7PayWvric5XABaFWQ30npWGL3DvvFw"

func Cliente_http(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/cliente_http/cliente_http.html", utilidades.Frontend))
	// Se conecta a la API
	cliente := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", Token)
	reg, err2 := cliente.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	defer reg.Body.Close()
	fmt.Println(reg.Status)

	// Leer cuerpo
	body, err := io.ReadAll(reg.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// La API devuelve probablemente un arreglo; usar slice
	var datos []modelos.Categoria
	errJson := json.Unmarshal(body, &datos)
	if errJson != nil {
		fmt.Println(errJson)
		return
	}

	// Retorno
	data := map[string][]modelos.Categoria{
		"datos": datos,
	}

	template.Execute(response, data)
}

func Cliente_http_crear(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/cliente_http/cliente_http_crear.html", utilidades.Frontend))

	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
	template.Execute(response, data)
}

func Cliente_http_crear_post(response http.ResponseWriter, request *http.Request) {
	mensaje := ""
	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El campo nombre es obligatorio. "
	}
	if mensaje != "" {
		utilidades.CrearMensajesFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/cliente-http/crear", http.StatusSeeOther)
	}
	//fmt.Fprintln(response,	"Nombre: " + request.FormValue("nombre"))

	datos := map[string]string{"nombre": request.FormValue("nombre")}
	jsonValue, _ := json.Marshal(datos)

	cliente := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.api.tamila.cl/api/categorias", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", Token)
	reg, err2 := cliente.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer reg.Body.Close()

	utilidades.CrearMensajesFlash(response, request, "success", "Se creo el cliente")
	http.Redirect(response, request, "/cliente-http/cliente-http-crear", http.StatusSeeOther)
}

func Cliente_http_editar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/cliente_http/cliente_http_editar.html", utilidades.Frontend))

	vars := mux.Vars(request)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", Token)

	reg, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reg.Body.Close()
	body, err := io.ReadAll(reg.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// CORRECCIÓN: Usar un objeto único, no un slice
	var datos modelos.Categoria
	errJson := json.Unmarshal(body, &datos)
	if errJson != nil {
		fmt.Println(errJson)
	}

	data := map[string]string{
		"id":     vars["id"],
		"nombre": datos.Nombre,
		"slug":   datos.Slug,
	}
	template.Execute(response, data)
}


func Cliente_http_editar_post(response http.ResponseWriter, request *http.Request) {
	mensaje := ""
	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El campo Nombre está vacío. "
	}
	if mensaje != "" {
	}

	vars := mux.Vars(request)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", Token)

	reg, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer reg.Body.Close()
	body, err := io.ReadAll(reg.Body)
	if err != nil {
		fmt.Println(err)
	}

	// La API a veces devuelve un objeto y a veces un array con un objeto.
	// Este código maneja ambos casos. No usamos el resultado,
	// pero evitamos el error de unmarshal en el log.
	if len(body) > 0 && body[0] == '[' {
		var a []modelos.Categoria
		json.Unmarshal(body, &a)
	} else {
		var c modelos.Categoria
		json.Unmarshal(body, &c)
	}
	//Se edita el registro
	datosJson := map[string]string{"nombre": request.FormValue("nombre")}
	jsonValue, _ := json.Marshal(datosJson)
	req2, err2 := http.NewRequest("PUT", "https://www.api.tamila.cl/api/categorias/"+vars["id"], bytes.NewBuffer(jsonValue))
	req2.Header.Set("Authorization", Token)
	if err2 != nil {
		fmt.Println(err2)
	}
	reg2, err3 := client.Do(req2)
	defer reg.Body.Close()
	if err3 != nil {
		fmt.Println(err3)
	}

	defer reg2.Body.Close()

	http.Redirect(response, request, "/cliente-http/cliente-http-editar/"+vars["id"], http.StatusSeeOther)
}

func Cliente_http_eliminar(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", Token)

	reg, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reg.Body.Close()
	body, err := io.ReadAll(reg.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var datos modelos.Categoria
	errJson := json.Unmarshal(body, &datos)
	if errJson != nil {
		fmt.Println(errJson)
	}

	// Se elimina el registro
	req2, err2 := http.NewRequest("DELETE", "https://www.api.tamila.cl/api/categorias/"+vars["id"], nil)
	req2.Header.Set("Authorization", Token)
	if err2 != nil {
		fmt.Println(err2)
	}
	reg2, err3 := client.Do(req2)
	defer reg.Body.Close()
	if err3 != nil {
		fmt.Println(err3)
	}

	defer reg2.Body.Close()
	http.Redirect(response, request, "/cliente-http", http.StatusSeeOther)
}