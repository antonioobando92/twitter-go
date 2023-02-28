package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar el password */
func EncriptarPassword(pass string) (string, error) {
	// Es un algoritmo, basado en 2, elevado al costo, entre mas costo mayor seguridad, pero se demora mas tiempo en procesar,
	// lo que hace es encriptarlo 256 veces por cada una de las pasadas. Se recomienda minimo 6 para usuarios comun y admin (8)
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
