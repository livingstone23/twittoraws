package bd

import (
	"context"
	"fmt"

	"github.com/livingstone23/twittoraws/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {

	fmt.Println("Funcion Busco Perfil")
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	//Utilizacin de paquete para poder filtar y paginar
	opciones := options.Find()
	opciones.SetLimit(10)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 10)

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		return resultados, false
	}

	for cursor.Next(ctx) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true

}