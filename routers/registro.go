package routers

import (
	"encoding/json"
	"net/http"

	"github.com/antonioobando92/twitter-go.git/bd"
	"github.com/antonioobando92/twitter-go.git/models"
)

/*Registro es la función para crear en la BD el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	// El body de httpRequest es un string, osea es un dato que se puede leer una sola vez y luego se destruye,
	// no se puede usar el body en varios lugares
	err := json.NewDecoder(r.Body).Decode(&t) // Lo guarda en una estructura de tipo JSON en el modelo Usuario.
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, estatus, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if estatus == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
