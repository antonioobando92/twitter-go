package jwt

import (
	"time"

	"github.com/antonioobando92/twitter-go.git/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook") // Clave privada, jwt trabaja con un slite de string
	payload := jwt.MapClaims{                                 // Lista de privilegios que se van a configurar en el Payload, segun el modelo de usuarios.
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicación,
		"sitioWeb":         t.SitioWeb,
		"_id":              t.ID.Hex(),                            // Devuelve el texto mostrando bien en formato hexadecimal
		"exp":              time.Now().Add(time.Hour * 24).Unix(), // Unix, lo devuelve en formato long, numerico y es bastante rapido
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave) // SignedString es el string de firma por seguridad
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
