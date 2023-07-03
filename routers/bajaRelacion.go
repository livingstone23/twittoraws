package routers

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func BajaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	fmt.Println("Ingresando a Reuter.BajaRelacion")

	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro del id es obligatorio."
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		r.Message = "Ocurrio un error al intentar Borrar la relacion" + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado borrar la relacion" + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Baja Relacion  OK"
	return r

}
