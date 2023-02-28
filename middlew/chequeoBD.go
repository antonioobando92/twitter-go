package middlew

import (
	"net/http"

	"github.com/antonioobando92/twitter-go.git/bd"
)

/* Es el middlew que me permite permite conocer el estado de la BD */
// Es un pasa manos, ingresa como parametro la función http.HandlerFunc y es lo que retorna, luego de hacer la comprobación de la BD
// Si la validación no es exitosa, retorna un error 500 y mata todo el endpoint y lo finaliza como tal.
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
