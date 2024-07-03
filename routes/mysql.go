package routes

import (
	"fmt"
	"net/http"
	"proyecto/connect"
	"proyecto/modelos"
	utils "proyecto/utilities"
	"text/template"
)

func Mysql_Listar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/mysql.html", utils.Frontend))

	// conexion a la base de datos
	connect.ConnectToDB()

	defer connect.CloseConnection()
	sql := "SELECT id, nombre, correo, telefono, password FROM clientes ORDER BY id DESC"
	clientes := modelos.Clients{}
	// datos, err := connect.MySqlDatabase.Query(sql)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	datos, err := connect.MySqlDatabase.Query(sql)
	if err != nil {
		fmt.Println("Error querying database:", err)
		data := map[string]any{
			"title": "MySQL Page",
			"users": modelos.Clients{},
		}
		template.Execute(response, data)
		return
	}
	defer datos.Close()

	for datos.Next() {
		cliente := modelos.Client{}
		err := datos.Scan(&cliente.ID, &cliente.Nombre, &cliente.Correo, &cliente.Telefono, &cliente.Password)
		if err != nil {
			fmt.Println(err)
		}
		clientes = append(clientes, cliente)
	}
	// retorno
	data := map[string]any{
		"title": "MySQL Page",
		"users": clientes,
	}

	template.Execute(response, data)

}

func CrearCliente(response http.ResponseWriter, request *http.Request) {

	template := template.Must(template.ParseFiles("views/crear-cliente.html", utils.Frontend))
	css, message := utils.ReturnAlertFlash(response, request)
	fmt.Println(message)
	data := map[string]any{
		"title":   "MySQL Page",
		"css":     css,
		"message": message,
	}
	template.Execute(response, data)
}

func CrearClientePOST(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	email := request.FormValue("email")
	// id := mux.Vars(request)["id"]

	if len(name) == 0 {
		message := "El nombre es requerido"
		utils.CreateAlertFlash(response, request, "danger", message)
		http.Redirect(response, request, "/crear-cliente", http.StatusSeeOther)
		return
	}
	if len(email) == 0 {
		message := "El email es requerido"
		utils.CreateAlertFlash(response, request, "danger", message)
		http.Redirect(response, request, "/crear-cliente", http.StatusSeeOther)
		return
	}

	connect.ConnectToDB()

	defer connect.CloseConnection()

	sql := fmt.Sprintf("INSERT INTO clientes (nombre, correo) VALUES ('%s', '%s')", name, email)

	_, err := connect.MySqlDatabase.Exec(sql)

	if err != nil {
		fmt.Println(err)
		utils.CreateAlertFlash(response, request, "danger", "Error al insertar")
		http.Redirect(response, request, "/crear-cliente", http.StatusSeeOther)
		return
	}

	http.Redirect(response, request, "/mysql", http.StatusFound)
}
