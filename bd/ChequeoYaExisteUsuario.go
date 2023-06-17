package bd

import (
	"context"
	"fmt"

	"github.com/livingstone23/twittoraws/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario: Funcion que recibe un email y chequea si el usuario ya existe*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	fmt.Println("Funcion ChequeoYaExisteUsuario ")
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")
	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
