package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/antonioobando92/twitter-go.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos lee los usuarios registrados, si se recibe "R" en quienes trae solo los que se relacionan conmigo */
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario // Enviar al http un slice de Usuarios

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20) // Solo devuelvo 20 registros

	// Expresion regular para comparaciones y busquedas de strings. "i" indica que no va a fijar si es mayuscula o miniscula.
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions) // El resultado de la consulta devuelve en un cursor
	if err != nil {
		fmt.Println("Error en la busqueda " + err.Error())
		return results, false
	}

	var encontrado, incluir bool
	for cur.Next(ctx) { // Recorre el cursor
		var s models.Usuario  // Se lee el modelo de usuario en cada uno en particular del cursor.
		err := cur.Decode(&s) // Se graba en el punturo del modelo usuario.
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex() // El ID del cursor

		incluir = false

		encontrado, _ = ConsultoRelacion(r) // se omite el bd. dado que esta en la misma carpeta la funcion.
		if tipo == "new" && !encontrado {
			incluir = true
		}

		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.UsuarioRelacionID == ID { // valida si el usuario, tiene relacion con el mismo usuario.
			incluir = false
		}

		// Se eliminan los valores de los campos que no se requieren devolver.
		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx) // Cierra el cursor del contexto.
	return results, true
}
