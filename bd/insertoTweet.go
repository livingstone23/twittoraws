package bd

import (
	"context"
	"fmt"

	"github.com/livingstone23/twittoraws/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	fmt.Println("Ingresando en funcion InsertoTweet")
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweet")

	//Creamos el registro a Insertar, lo vamos a convertir a bson
	registro := bson.M{
		"userid":  t.UserId,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
