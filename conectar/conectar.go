package conectar

import (
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Funcion para conectar a la Base de datos
var Db *sql.DB

func Conectar() {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	conection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	Db = conection

}

// Cerrar Conexion
func CerrarConexion() {
	Db.Close()
}
