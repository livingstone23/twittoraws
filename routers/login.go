package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/jwt"
	"github.com/livingstone23/twittoraws/models"
)

/*Login permite confirmar si el usuario esta acreditado*/
func Login(ctx context.Context) models.RespApi {

	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Usuario y/o Contraseña inválidos " + err.Error()
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "El email del usuario es requerido "
		return r
	}

	userData, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		r.Message = "Usuario y/o Contraseña inválidos "
		return r
	}

	//modulo para llamar a rutina que genera token
	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Ocurrio un error al instante de generar un Token correspondiente" + err.Error()
		return r
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Ocurrio un error al instante formatear el Token correspondiente" + err2.Error()
		return r
	}

	//como grabrar una cookie desde el backend

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(720 * time.Hour),
	}

	cookieString := cookie.String()

	//para grabar la cookie
	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res
	return r

}
