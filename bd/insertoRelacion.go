package bd

import (
	"context"
	"fmt"

	"github.com/livingstone23/twittoraws/models"
)

func InsertoRelacion(t models.Relacion) (bool, error) {
	fmt.Println("Ingresando en funcion InsertoRelacion")

	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relacion")

	//Creamos el registro a Insertar, lo vamos a convertir a bson

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
