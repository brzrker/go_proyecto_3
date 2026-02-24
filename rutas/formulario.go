package rutas

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"proyecto_3/utilidades"
	"proyecto_3/validaciones"
	"strings"
	"time"
)

func Formulario_get(response http.ResponseWriter, request *http.Request) { 
	
	template := template.Must(template.ParseFiles("templates/formularios/formulario.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css" :		css_sesion,
		"mensaje" : css_mensaje,
	}
	template.Execute(response, data)

}

func Formulario_post(response http.ResponseWriter, request *http.Request) {
	mensaje := ""
	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El campo nombre es obligatorio. "
	}
	if len(request.FormValue("correo")) == 0 {
		mensaje = mensaje + "El campo correo es obligatorio. "
	}
	if validaciones.Regex_correo.FindStringSubmatch(request.FormValue("correo")) == nil {
		mensaje = mensaje + "El correo no es valido. "
	}
	if validaciones.ValidarPassword(request.FormValue("password")) == false {
		mensaje = mensaje + "La contraseña no es valida. "
	}
	if mensaje != "" {
		//fmt.Fprintln(response, mensaje)
		//return
		utilidades.CrearMensajesFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/formulario", http.StatusSeeOther)
		return
	}
	//p2gHNiENUw CONTRASEÑA VALIDA PARA PRUEBAS
	fmt.Fprintln(response,	"Nombre: " + request.FormValue("nombre") + // request.FormValue("nombre") es el name del input del formulario
							" | Correo: " + request.FormValue("correo") + 
							" | Telefono: " + request.FormValue("telefono") + 
							" | Password: " + request.FormValue("password"))
}

func Formulario_upload(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/formularios/upload.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css" :		css_sesion,
		"mensaje" : css_mensaje,
	}
	template.Execute(response, data)
}

func Formulario_upload_post(response http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("foto")
	if err != nil {
		utilidades.CrearMensajesFlash(response, request, "danger", "Ocurrio un error al subir la foto.1")
	}

	var extencion = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ") // time[0] = fecha, time[1] = hora
	foto := string(time[4][6:14]) + "." + extencion // time[4][6:14] = hora sin los : para que no de error al guardar la foto
	var archivo string = "static/upload/foto/" + foto
	f, errCopy := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if errCopy != nil {
		utilidades.CrearMensajesFlash(response, request, "danger", "Ocurrio un error al subir la foto.2")
	}

	_, errCopiar := io.Copy(f, file)
	if errCopiar != nil {
		utilidades.CrearMensajesFlash(response, request, "danger", "Ocurrio un error al subir la foto.3")
	}
	// Aca se guardaran en la base de datos

	// redireccionamos al formulario con un mensaje de exito
	utilidades.CrearMensajesFlash(response, request, "success", "Foto " + foto + " subida correctamente.")
	http.Redirect(response, request, "/formulario/upload", http.StatusSeeOther)
}



/* func Formulario_post(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response,	"Nombre: " + request.FormValue("nombre") + // request.FormValue("nombre") es el name del input del formulario
							" | Correo: " + request.FormValue("correo") + 
							" | Telefono: " + request.FormValue("telefono") + 
							" | Password: " + request.FormValue("password"))
} */