package rutas

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"proyecto_3/conectar"
	"proyecto_3/modelos"
	"proyecto_3/utilidades"
	"proyecto_3/validaciones"
	"github.com/gorilla/mux"
)

func My_SQL_listar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/my_sql/my_sql.html", utilidades.Frontend))

	// Coneccion a la base de datos
	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes order by id desc"
	clientes := modelos.Clientes{}
	datos, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()
	for datos.Next() {
		dato := modelos.Cliente{}
		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		clientes = append(clientes, dato)
	}

	// Retorno
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := modelos.ClienteHttp {
		Css : css_sesion,
		Mensaje : css_mensaje,
		Datos : clientes,
	}
	template.Execute(response, data)
}

func My_SQL_crear(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/my_sql/my_sql_crear.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
	template.Execute(response, data)
}

func My_SQL_crear_post(response http.ResponseWriter, request *http.Request) {
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
	if mensaje != "" {
		utilidades.CrearMensajesFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/my-sql/my-sql-crear", http.StatusSeeOther)
		return
	}

	conectar.Conectar()

	sql := "INSERT INTO clientes (nombre, correo, telefono) VALUES (?,?,?)"
	_, err := conectar.Db.Exec(sql, request.FormValue("nombre"), request.FormValue("correo"), request.FormValue("telefono"), )
	if err != nil {
		fmt.Println(err)
	}

	utilidades.CrearMensajesFlash(response, request, "success", "Se creo el cliente")
	http.Redirect(response, request, "/my-sql/my-sql-crear", http.StatusSeeOther)
}

func My_SQL_editar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/my_sql/my_sql_editar.html", utilidades.Frontend))
	
	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes WHERE id = ?"

	vars := mux.Vars(request)
	datos, err := conectar.Db.Query(sql, vars["id"])
		if err != nil {
		fmt.Println(err)
		return
	}

	defer conectar.CerrarConexion()
	var dato modelos.Cliente
	for datos.Next() {
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			log.Fatal(err)
		}
	}

	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	clienteHttp := modelos.ClienteHttp2{
		Css: css_sesion,
		Mensaje: css_mensaje,
		Datos: dato,
	}
	template.Execute(response, clienteHttp)
}

func My_SQL_editar_post(response http.ResponseWriter, request *http.Request) {
	mensaje := ""
	vars := mux.Vars(request)
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
	if mensaje != "" {
		utilidades.CrearMensajesFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/my-sql/my-sql-editar/"+vars["id"], http.StatusSeeOther)
		return
	}

	conectar.Conectar()

	sql := "UPDATE clientes SET nombre = ?, correo = ?, telefono = ? WHERE id = ?;"
	_, err := conectar.Db.Exec(sql, request.FormValue("nombre"), request.FormValue("correo"), request.FormValue("telefono"), vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}

	utilidades.CrearMensajesFlash(response, request, "success", "Se modificó el cliente")
	http.Redirect(response, request, "/my-sql/my-sql-editar/"+vars["id"], http.StatusSeeOther)
}

func My_SQL_eliminar(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	conectar.Conectar()

	sql := "DELETE FROM clientes WHERE id = ?;"
	_, err := conectar.Db.Exec(sql, vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}

	utilidades.CrearMensajesFlash(response, request, "success", "Se Eliminó el cliente")
	http.Redirect(response, request, "/my-sql", http.StatusSeeOther)
}