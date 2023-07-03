package routers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func AltaRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	fmt.Println("Ingresando a Reuter.AltaRelacion")

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

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		r.Message = "Ocurrio un error al intentar insertar la relacion" + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar la relacion" + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Alta de relacion OK"
	return r

}
