package bd

import (
	"context"
	"log"
	"time"

	"github.com/antonioobando92/twitter-go.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweets lee los tweets de un perfil */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()                          // Opciones para realizar la consulta hacia mongo.
	opciones.SetLimit(20)                               // Cuantos documentos quiere que vaya retornando por pagina
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) // Ordenarlo por fecha de forma descendente.
	opciones.SetSkip((pagina - 1) * 20)                 // Saltear la cantidad de registros segun la pagina, primero 20, luego 40, asi consecutivamente.

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	// context.TODO() crea un contexto nuevo vacio, diferente al ctx que solo esta configurado para 15 seg.
	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
