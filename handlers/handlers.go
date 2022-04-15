package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/fdymendo/EjemploTwitter/middlew"
	"github.com/fdymendo/EjemploTwitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores de mi puerto, el handler y pongo a escuchar al server*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
