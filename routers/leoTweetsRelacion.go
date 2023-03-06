package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/antonioobando92/twitter-go.git/bd"
)

/*LeoTweetsSeguidores lee los tweets de todos nuestros seguidores */
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina: ", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) // strconv.Atoi conversion de string a entero
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a 0: ", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets: ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)    // Devuelve estatus 200 ó 201
	json.NewEncoder(w).Encode(respuesta) // Retorna la respuesta.

}
