package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/antonioobando92/twitter-go.git/bd"
	"github.com/antonioobando92/twitter-go.git/models"
)

/*SubirAvatar sube la imagen al servidor */
func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, _ := r.FormFile("avatar")                             // Procesa como si fuera un formulario de HTML y el nombre del archivo avatar
	var extension = strings.Split(handler.Filename, ".")[1]              // Devuelve un String [1] porque sino devolveria un verctor, hace un split del .
	var archivo string = "uploads/avatars" + IDUsuario + "." + extension // La ruta donde va a guardar el archivo, le agrega el id del usuario y la extensión del archivo

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666) // 0666 permisos de lectura, escritura y ejecucion
	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar en la BD! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
