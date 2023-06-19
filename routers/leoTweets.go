package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func LeoTweets(request events.APIGatewayProxyRequest) models.RespApi {

	var r models.RespApi
	r.Status = 400
	fmt.Println("Ingresando a funcion LeoTweet")

	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio "
		return r
	}

	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar el parametro de pagina como un valor mayor a 0"
		return r
	}

	tweets, correcto := bd.LeoTweets(ID, int64(pag))
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
