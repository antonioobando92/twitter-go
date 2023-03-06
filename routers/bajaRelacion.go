package routers

import (
	"net/http"

	"github.com/antonioobando92/twitter-go.git/bd"
	"github.com/antonioobando92/twitter-go.git/models"
)

/*BajaRelacion realiza el borrado de la relación entre usuarios */
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio ", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se la logrado borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
