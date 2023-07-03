package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func ListaUsuarios(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	fmt.Println("Ingresando a Reuter.ListaUsuarios")

	var r models.RespApi
	r.Status = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUsuario := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Debe enviar el parametro 'page' como entero o mayor a 0 " + err.Error()
		return r
	}

	usuarios, status := bd.LeoUsuariosTodos(IDUsuario, int64(pageTemp), search, typeUser)
	if !status {
		r.Message = "Error al leer a los usuarios"
		return r
	}

	respJson, err := json.Marshal(usuarios)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios en JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r

}
