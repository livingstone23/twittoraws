package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.RespApi {

	var r models.RespApi
	r.Status = 400
	fmt.Println("Ingresando en funcion ModificarPerfil")

	var t models.Usuario
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = "Datos incorrectos " + err.Error()
		fmt.Println(r.Message)
	}

	status, err := bd.ModificoRegistro(t, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrio un error al intentar modificar el registro. " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado modificar el registro del usuario. "
		return r
	}

	r.Status = 200
	r.Message = "Modificacion perfil OK"
	fmt.Println(r.Message)
	return r

}
