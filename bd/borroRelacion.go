package bd

import (
	"context"
	"fmt"

	"github.com/livingstone23/twittoraws/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	fmt.Println("Ingresando en funcion bd.BorroRelacion")

	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relacion")

	//Creamos el registro a Insertar, lo vamos a convertir a bson

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
