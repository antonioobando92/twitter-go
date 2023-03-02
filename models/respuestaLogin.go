package models

/*RespuestaLogin tiene el token que se devuelve con el login */
type RespuestaLogin struct {
	// Se le indica que devuelve un JSON, se va a llamar token en minuscula
	// y omitempty por si hay algun error devuelve la estructura vacio.
	Token string `json:"token,omitempty"`
}
