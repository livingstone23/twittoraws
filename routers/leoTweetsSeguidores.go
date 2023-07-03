package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func LeoTweetsSeguidores(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	fmt.Println("Ingresando a Reuter.LeoTweetsSeguidores")

	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	pagina := request.QueryStringParameters["pagina"]
	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar el parametro de pagina como un valor mayor a 0"
		return r
	}

	tweets, correcto := bd.LeoTweetsSeguidores(IDUsuario, pag)
	if !correcto {
		r.Message = "Error al leer los Tweets "
		return r
	}

	resJson, err := json.Marshal(tweets)
	if err != nil {
		r.Message = "Error al formatear los datos de los usuarios de JSON"
		return r
	}

	r.Status = 200
	r.Message = string(resJson)
	return r

}
