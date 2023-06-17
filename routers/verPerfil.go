package routers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.RespApi {

	var r models.RespApi
	r.Status = 400
	fmt.Println("Ingresando en funcion VerPerfil")

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio "
		return r
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Ocurrio un error al intentar buscar el registro "
		return r
	}

	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios como JSON" + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r

}
