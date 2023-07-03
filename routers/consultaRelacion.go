package routers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func ConsultaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	fmt.Println("Ingresando a Reuter.ConsultaRelacion")

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

	var resp models.RespuestaConsultaRelacion

	hayRelacion := bd.ConsultoRelacion(t)
	if !hayRelacion {
		resp.Status = false

	} else {
		resp.Status = true
	}

	respJson, err := json.Marshal(hayRelacion)

	if err != nil {
		r.Message = "Error al formatear los datos de los usuarios como JSON" + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r

}
