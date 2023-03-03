package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/antonioobando92/twitter-go.git/middlew"
	"github.com/antonioobando92/twitter-go.git/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores, seteo mi puerto, el handler y pongo a escuchar el servidor */
func Manejadores() {
	router := mux.NewRouter()

	// HandleFunc, es una función de gorilla/mux
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT") // Revisa en el OS si ya esta creado en la variable de entorno
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(": "+PORT, handler))
}
