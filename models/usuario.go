package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Usuario es el modelo de usuario de la base de MongoDB */
// bson: Dato de entrada a la BD.
// json: Dato de salida en formato JSON como respuesta.
// ObjectID (Se grava en binario _id), bson omitempty indica que si llega algun dato viene vacio lo omita,
// que no lo tenga en cuenta para formar ningun JSON. json: id el formato en que va a devolver el JSON.
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`        // ObjectID (Se grava en binario _id), bson omitempty indica que si llega algun dato viene vacio lo omita, que no lo tenga en cuenta para formar ningun JSON. json: id el formato en que va a devolver el JSON.
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"` // Que no devuelva el dato en el JSON si no lo encuentra, porque el nombre lo va a tener siempre en la BD.
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`                 // Cuando se loquea es con email y cuando devuelva siempre devuelve el email.
	Password        string             `bson:"password" json:"password,omitempty"` // No se devuelve un password en el navegador
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicación       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
}
