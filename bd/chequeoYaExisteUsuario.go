package bd

import (
	"context"
	"time"

	"github.com/antonioobando92/twitter-go.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parámetro y chequea si ya esta en la BD */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Restricción timeout 15 seg, context.Background() actual de la BD
	defer cancel()                                                           // Se ejecuta como ultima instrucción de la función cancelando el contexto que se creo dentro del contexto backgroud, de esta forma limpia lo que se creo.

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email} // Buscar en la colección por el nombre de la variable, devuelve un formato bson.

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado) // Realiza la busqueda de 1 solo registro y lo decodifica en formato JSON, lo guarda en el modelo resultado
	ID := resultado.ID.Hex()                              // convierte el ObjectID a un Hexadecimal en formato String. Si resultado esta vacio, guarda ID vacio, no genera ningun tipo de error al tratar de acceder al dato y no exista.
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
