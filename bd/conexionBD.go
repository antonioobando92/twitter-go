package bd // El trata a todos los archivo deltro del paquete como un solo archivo, por lo cual en los archivos no se pueden
// tener funciones con el mismo nombre

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la BD */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:antonio1192@clustermongodb.j2sxkil.mongodb.net/?retryWrites=true&w=majority")

/* ConectarBD es la función que me permite conectar la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) // Hace un llamado para ver si la bd esta arriba
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/* ChequeoConnection es el ping a la BD */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil) // Hace un llamado para ver si la bd esta arriba
	if err != nil {
		return 0
	}
	return 1
}
