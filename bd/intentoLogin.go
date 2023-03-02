package bd

import (
	"github.com/antonioobando92/twitter-go.git/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la BD */
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email) // Como la función se encuentra en el mismo paquete no es necesario, definir el paquete
	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)                               // Password enviada por el cliente sin encriptar
	passwordBD := []byte(usu.Password)                              // Password Encriptada retornada desde la BD
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes) // Compara las dos password, si devuelve un error es porque no coincidio
	if err != nil {
		return usu, false
	}
	return usu, true
}
