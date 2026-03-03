package rutas

import (
	"fmt"
	"html/template"
	"net/http"
	"proyecto_3/conectar"
	"proyecto_3/modelos"
	"proyecto_3/utilidades"
	"proyecto_3/validaciones"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Seguridad_registro(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/seguridad/registro.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
	template.Execute(response, data)
}

func Seguridad_registro_post(response http.ResponseWriter, request *http.Request) {
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
	if len(request.FormValue("telefono")) == 0 {
		mensaje = mensaje + "El campo telefono es obligatorio. "
	}
	if validaciones.ValidarPassword(request.FormValue("password")) == false {
		mensaje = mensaje + "La contraseña no es valida. "
	}
	if mensaje != "" {
		utilidades.CrearMensajesFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/seguridad/registro", http.StatusSeeOther)
	}

	conectar.Conectar()
	sql := "INSERT INTO usuarios (nombre, correo, telefono, password) VALUES (?,?,?,?);"
	// Se genera la contraseña en hash
	//p2gHNiENUw
	costo := 8
	bytes, _ := bcrypt.GenerateFromPassword([]byte(request.FormValue("password")), costo)
	_, err := conectar.Db.Exec(sql, request.FormValue("nombre"), request.FormValue("correo"), request.FormValue("telefono"), string(bytes))
	if err != nil {
		fmt.Println(err)
	}

	utilidades.CrearMensajesFlash(response, request, "success", "Se creo el cliente")
	http.Redirect(response, request, "/seguridad/registro", http.StatusSeeOther)
}

func Seguridad_login(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/seguridad/login.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
	template.Execute(response, data)
}

func Seguridad_login_post(response http.ResponseWriter, request *http.Request) {
	mensaje := ""
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
		utilidades.CrearMensajesFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
	}

	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono, password FROM usuarios WHERE correo = ?"
	datos, err := conectar.Db.Query(sql, request.FormValue("correo"))
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()
	var dato modelos.Usuario
	for datos.Next() {
		errNext := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono, &dato.Password)
		if errNext != nil {
			utilidades.CrearMensajesFlash(response, request, "danger", "Las credenciales son incorrectas")
			http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
		}
	}
	
	// Se comparan los hash
	passwordBytes := []byte(request.FormValue("password"))
	passwordBD := []byte(dato.Password)
	errPassword := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if errPassword != nil {
		utilidades.CrearMensajesFlash(response, request, "danger", "La contraseña es incorrecta")
		http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
	} else {
		session, _ := utilidades.Store.Get(request, "session-name")
		id_string := strconv.Itoa(dato.Id) // Convierte el id en string porque en int da problemas
		session.Values["hola_id"] = id_string
		session.Values["hola_nombre"] = dato.Nombre
		err2 := session.Save(request, response)
		if err2 != nil {
			http.Error(response, err2.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(response, request, "/seguridad/protegida", http.StatusSeeOther)
	}
}

func Seguridad_protegida(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/seguridad/protegida.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	hola_id, hola_nombre := utilidades.RetornarLogin(request)
	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
		"hola_id": hola_id,
		"hola_nombre": hola_nombre,
	}
	template.Execute(response, data)
}

func Seguridad_logout(response http.ResponseWriter, request *http.Request) {
	session, _ := utilidades.Store.Get(request, "session-name")
	session.Values["hola_id"] = nil
	session.Values["hola_nombre"] = nil
	err2 := session.Save(request, response)
	if err2 != nil {
		http.Error(response, err2.Error(), http.StatusInternalServerError)
		return
	}
	utilidades.CrearMensajesFlash(response, request, "primary", "Se ha cerrado tu sesión")
	http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
}