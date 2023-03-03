package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/antonioobando92/twitter-go.git/bd"
)

/*LeoTweets leo los tweets de la base de datos */
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id: ", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina: ", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) // strconv.Atoi conversion de string a entero
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a 0: ", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer los tweets: ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Devuelve estatus 200 ó 201
	json.NewEncoder(w).Encode(respuesta)

}
