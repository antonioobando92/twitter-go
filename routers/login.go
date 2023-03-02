package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/antonioobando92/twitter-go.git/bd"
	"github.com/antonioobando92/twitter-go.git/jwt"
	"github.com/antonioobando92/twitter-go.git/models"
)

/*Login realiza el login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El Email del usuario es requerido ", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña invalidos ", 400)
		return
	}

	// JWT, manera de generar token o credenciales que se van a poder utilizar en todo el largo de la aplicación.
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Devuelve estatus 200 ó 201
	json.NewEncoder(w).Encode(resp)

	// Almacenar en la cookie de la maquina del usuario
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
