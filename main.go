package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proyecto/routes"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	mux := mux.NewRouter()

	//* rutas
	mux.HandleFunc("/", routes.Home)
	mux.HandleFunc("/about", routes.About)
	mux.HandleFunc("/params/{id:.*}/{slug:.*}", routes.Params)
	mux.HandleFunc("/querystring", routes.Querystring)
	mux.HandleFunc("/structures", routes.Structures)
	mux.HandleFunc("/form", routes.Formularios)
	mux.HandleFunc("/form-post", routes.FormularioPost).Methods("POST")
	mux.HandleFunc("/crear-cliente", routes.CrearCliente)
	mux.HandleFunc("/crear-cliente-post", routes.CrearClientePOST).Methods("POST")

	mux.HandleFunc("/upload", routes.FileUpload)
	mux.HandleFunc("/upload-post", routes.FileUploadSave).Methods("POST")

	mux.HandleFunc("/util-resources", routes.Resources)
	mux.HandleFunc("/util-resources/pdf", routes.PdfMaker)

	mux.HandleFunc("/api-consumer", routes.ApiConsumer)
	mux.HandleFunc("/mysql", routes.Mysql_Listar)

	//* Not found handler
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(routes.Page404).GetHandler()

	//* archivos estaticos
	s := http.StripPrefix("/public/", http.FileServer(http.FS(os.DirFS("./public/"))))
	mux.PathPrefix("/public/").Handler(s)

	//* variables de entorno
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	//* servidor
	server := &http.Server{
		Addr:         "127.0.0.1:" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on port " + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}

// func main() {
// 	// mux := http.NewServeMux()
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Listening on port 8081")
// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }

// func handler(response http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintln(response, "AAAAAAAaaaaaaaa!")
// }
