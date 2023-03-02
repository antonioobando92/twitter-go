package routers

import (
	"errors"
	"strings"

	"github.com/antonioobando92/twitter-go.git/bd"
	"github.com/antonioobando92/twitter-go.git/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor de Email usado en todos los endpoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los endpoints */
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) { // el parametro error, se debe colocar al final.
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{} // Se rquiere que el modelo, este marcado como un puntero.

	splitToken := strings.Split(tk, "Bearer") // Se convierte en un vector y en la posicion 0, va e star bearer y posicion 1 token.
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido") // No se debe enviar mayusculas, ni caracteres especiales.
	}
	tk = strings.TrimSpace(splitToken[1]) // Quita los espacios del token.
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
