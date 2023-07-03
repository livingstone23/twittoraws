package bd

import (
	"context"
	"fmt"

	"github.com/livingstone23/twittoraws/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) bool {
	fmt.Println("Ingresando en funcion bd.ConsultaRelacion")

	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false
	}
	return true

}
