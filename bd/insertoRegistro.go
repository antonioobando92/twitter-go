package bd

import (
	"context"
	"time"

	"github.com/antonioobando92/twitter-go.git/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistro es la parada final con la BD para insertar los datos del usuario */
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Restricción timeout 15 seg, context.Background() actual de la BD
	defer cancel()                                                           // Se ejecuta como ultima instrucción de la función cancelando el contexto que se creo dentro del contexto backgroud, de esta forma limpia lo que se creo.

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	// Es la forma de obtener el id que se creo en mongo y lo retorna en tipo string
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
