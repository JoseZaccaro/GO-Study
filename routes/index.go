package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"proyecto/modelos"
	utils "proyecto/utilities"
	"proyecto/validaciones"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func Index() {

}

func Home(response http.ResponseWriter, request *http.Request) {

	template := template.Must(template.ParseFiles("views/index.html", utils.Frontend))
	data := map[string]string{
		"title": "Hola mundo",
	}
	template.Execute(response, data) // template, err := template.ParseFiles("views/index.html", "layout/front.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data := map[string]string{
	// 	"title": "Hola mundo",
	// }
	// template.Execute(response, data)
	// fmt.Fprintln(response, "Home page route")
}

func About(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/about.html", utils.Frontend))
	// template, err := template.ParseFiles("views/about.html", "layout/front.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	data := map[string]string{
		"title": "About",
	}
	template.Execute(response, data)

	// fmt.Fprintln(response, "About page route")
}

func Params(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	data := map[string]string{
		"id":    vars["id"],
		"slug":  vars["slug"],
		"texto": "hola mundo",
		"title": "Params",
	}

	template, err := template.ParseFiles("views/params.html", "layout/front.html")

	if err != nil {
		log.Fatal(err)
	}
	template.Execute(response, data)
	// fmt.Fprintf(response, "Params page route with id %s and slug %s", vars["id"], vars["slug"])
}

func Querystring(response http.ResponseWriter, request *http.Request) {

	params := request.URL.Query().Get("params")
	algo := request.URL.Query().Get("algo")

	data := map[string]string{
		"params": params,
		"algo":   algo,
		"title":  "Querystring",
	}

	template, err := template.ParseFiles("views/query.html", "layout/front.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(response, data)
	// fmt.Fprintln(response, "Querystring page route: ")
	// fmt.Fprintln(response, request.URL)
	// fmt.Fprintln(response, request.URL.RawQuery)
	// fmt.Fprintln(response, request.URL.Query())

	// fmt.Fprintf(response, "Params: %s, Algo: %s", params, algo)

}

type Ability struct {
	Nombre string
}

type Person struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Ability
}

func Structures(response http.ResponseWriter, request *http.Request) {

	template, err := template.ParseFiles("views/structures.html", "layout/front.html")
	// data := map[string]string{
	// 	"algo": "hola mundo",
	// }
	if err != nil {
		log.Fatal(err)
	}

	ability1 := Ability{"Java"}
	ability2 := Ability{"Python"}
	abilities := []Ability{ability1, ability2}
	person := Person{"Pedro", 23, 3, abilities}
	data := map[string]any{
		"title":  "Structures",
		"Person": person,
	}

	template.Execute(response, data)

}

func Page404(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/Page404.html", utils.Frontend))
	data := map[string]string{
		"title": "404",
	}
	template.Execute(response, data)
}

func Formularios(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/form.html", utils.Frontend))

	css, message := utils.ReturnAlertFlash(response, request)

	fmt.Println("css:", css)
	fmt.Println("message:", message)

	data := map[string]string{
		"title":   "Form Page",
		"css":     css,
		"message": message,
	}

	template.Execute(response, data)
}
func FormularioPost(response http.ResponseWriter, request *http.Request) {

	mensaje := ""
	name := request.FormValue("name")
	phone := request.FormValue("phone")
	email := request.FormValue("email")
	password := request.FormValue("password")

	if len(name) == 0 {
		mensaje = "Name is required. "
	}

	if len(phone) == 0 {
		mensaje = mensaje + "Phone is required. "
	}

	if len(email) == 0 {
		mensaje = mensaje + "Email is required. "
	}

	if len(password) == 0 {
		mensaje = mensaje + "Password is required. "
	}

	if validaciones.Regex_correo.FindStringSubmatch(email) == nil {
		mensaje = mensaje + "Email is not valid. "
	}

	if !validaciones.ValidarPassword(password) {
		mensaje = mensaje + "Password must have at least 6 characters and at least one number and one uppercase and one lowercase letter. "
	}

	if mensaje != "" {
		// fmt.Fprintln(response, mensaje)
		// return
		utils.CreateAlertFlash(response, request, "danger", mensaje)
		http.Redirect(response, request, "/form", http.StatusSeeOther)
	}

	fmt.Fprintf(response, "Name: %s, Phone: %s, Email: %s, Password: %s", name, phone, email, password)

	// fmt.Fprintln(response, request.FormValue("name"))
	// template := template.Must(template.ParseFiles("views/form-post.html", utils.Frontend))
	// data := map[string]string{
	// 	"title": "Form POST",
	// }
	// template.Execute(response, data)
}

func FileUpload(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/upload.html", utils.Frontend))

	// css, message := utils.ReturnAlertFlash(response, request)

	// fmt.Println("css:", css)
	// fmt.Println("message:", message)

	data := map[string]string{
		"title": "File Upload Page",
		// "css":     css,
		// "message": message,
	}

	template.Execute(response, data)
}

func FileUploadSave(response http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("photo")

	if err != nil {
		fmt.Fprintf(response, "%v", err)
		return
	}
	defer file.Close()

	var fileName, fileExtension string = strings.Split(handler.Filename, ".")[0], strings.Split(handler.Filename, ".")[1] // Obtener la extensión del archivo
	time := strings.Split(time.Now().String(), " ")
	photo := string(time[4][6:14]) + "." + fileExtension
	fmt.Printf(" File name: %s\n File extension: %s", fileName, fileExtension)
	var archivo string = "public/uploads/photo/" + photo
	f, errCopy := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)

	if errCopy != nil {
		utils.CreateAlertFlash(response, request, "danger", errCopy.Error())
		return
	}

	_, errCopiar := io.Copy(f, file)
	if errCopiar != nil {
		utils.CreateAlertFlash(response, request, "danger", errCopiar.Error())
		return
	}
	// Aca guardaria en caso de guardar en DB
	utils.CreateAlertFlash(response, request, "success", "File uploaded with success")
	http.Redirect(response, request, "/upload", http.StatusSeeOther)

}

func PdfMaker(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/pdf.html", utils.Frontend))

	data := map[string]string{
		"title": "Pdf maker Page",
	}

	template.Execute(response, data)
}
func Resources(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/utils.html", utils.Frontend))

	data := map[string]string{
		"title": "Utils Page",
	}

	template.Execute(response, data)
}

func ApiConsumer(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/api-consumer.html", utils.Frontend))
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://rickandmortyapi.com/api/character", nil)

	if err != nil {
		fmt.Println(err)
	}
	reg, _ := client.Do(req)
	model := modelos.Response{}
	body, _ := io.ReadAll(reg.Body)
	defer reg.Body.Close()
	errJSON := json.Unmarshal(body, &model)
	fmt.Println(reg.Status)
	if errJSON != nil {
		fmt.Println(errJSON)
	}

	// var jsonBody interface{}
	// err = json.Unmarshal(body, &jsonBody)
	// if err != nil {
	// 	fmt.Println("Error al deserializar el JSON:", err)
	// 	return
	// }

	// // Convertir el JSON deserializado de nuevo a una cadena JSON con sangría
	// prettyJSON, err := json.MarshalIndent(jsonBody, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error al formatear el JSON:", err)
	// 	return
	// }

	// fmt.Println(string(prettyJSON))

	data := map[string]any{
		"title": "Api consumer Page",
		"body":  model,
	}

	template.Execute(response, data)
}
