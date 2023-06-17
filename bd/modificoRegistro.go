package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/livingstone23/twittoraws/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Modifico registro permite modificar el perfil del usuario*/
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	fmt.Println("Ingresando ModificoRegistro")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittordb")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})

	fmt.Println("Modificando nombre por: " + u.Nombre)
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	fmt.Println("Modificando apellido por: " + u.Apellido)
	if len(u.Apellido) > 0 {
		registro["apellidos"] = u.Apellido
	}

	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}

	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}

	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}

	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}

	objId, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil

}
