package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.RespApi {

	var mensaje models.Tweet
	var r models.RespApi
	r.Status = 400
	fmt.Println("Ingresando en funcion GraboTweet")
	IDUsuario := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)

	if err != nil {
		r.Message = "Ocurrio un error al intentar decodificar el body" + err.Error()
		return r
	}

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		r.Message = "Ocurrio un error al intentar insertar un registro, reintente nuevamente" + err.Error()
		return r
	}

	if status == false {
		r.Message = "No se ha logrado insertar el tweet" + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Tweet creado correctamente. "
	return r

}
