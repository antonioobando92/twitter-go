package bd

import (
	"context"
	"time"

	"github.com/antonioobando92/twitter-go.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los tweets de mis seguidores */
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20        // Paginado de a 20 registros
	condiciones := make([]bson.M, 0) // indica que tiene 0 elementos inicialmente para no crear elementos de mas.
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{ // lookup, se usa para unir dos tablas (colecciones) en mongo relacion y tweet
		"$lookup": bson.M{
			"from":         "tweet",             // Con que tabla se quiere unir la tabla relacion, en este caso tweet
			"localField":   "usuariorelacionid", // indicar el campo por el cual se va a unir las 2 tablas, campo de Relacion
			"foreignField": "userid",            // indicar el campo de tweet
			"as":           "tweet",             // Indica un alias a la tabla tweet, en este caso se deja igual.
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})          // Si no usuaramos unwind los resultados devolverian los datos en un formato que no nos sirve, maestro - detalle, lo que permite es que todos los documentos viene igual, no viene maestro detalle, sino informacion repetida
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}}) // Para ordenar los datos de una manera determinada, por el campo fecha. -1 descendente (mas antiguo al final) y 1 de menor a mayor
	condiciones = append(condiciones, bson.M{"$skip": skip})                // Para paginar por la cantidad de registros configurado
	condiciones = append(condiciones, bson.M{"$limit": 20})                 // Indica cada cuanto salta en el paginador para mostrar los datos bien.

	cursor, err := col.Aggregate(ctx, condiciones) // Framework de mongo Aggregate, se ejecuta en la BD y crea un cursor
	var result []models.DevuelvoTweetsSeguidores   // Se crea una variable
	err = cursor.All(ctx, &result)                 // Se recorre el cursor de un solo tiron o vez y se envia a la variable result.
	if err != nil {
		return result, false
	}
	return result, true
}
