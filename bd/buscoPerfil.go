package bd

import (
	"context"
	"fmt"

	"github.com/livingstone23/twittoraws/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.Usuario, error) {

	fmt.Println("Funcion Busco Perfil")
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objId, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objId,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}
	return perfil, nil

}
