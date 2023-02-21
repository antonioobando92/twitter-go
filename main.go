package main

import (
	"log"

	"github.com/antonioobando92/twitter-go/bd"
	"github.com/antonioobando92/twitter-go/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD!")
		return
	}
	handlers.Manejadores()
}
