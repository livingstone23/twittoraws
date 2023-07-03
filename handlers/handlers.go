package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/twittoraws/jwt"
	"github.com/livingstone23/twittoraws/models"
	"github.com/livingstone23/twittoraws/routers"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi
	r.Status = 400

	isOK, statusCode, msg, claim := validoAuthorization(ctx, request)
	if !isOK {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	//Permite organizar nuestras rutas.
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)

		case "login":
			return routers.Login(ctx)

		case "tweet":
			return routers.GraboTweet(ctx, claim)
		case "altaRelacion":
			return routers.AltaRelacion(ctx, request, claim)
		case "subirAvatar":
			return routers.UploadImage(ctx, "A", request, claim)
		case "subirBanner":
			return routers.UploadImage(ctx, "B", request, claim)
		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "verperfil":
			return routers.VerPerfil(request)
		case "leotweets":
			return routers.LeoTweets(request)
		case "obtenerAvatar":
			return routers.ObtenerImagen(ctx, "A", request, claim)
		case "obtenerBanner":
			return routers.ObtenerImagen(ctx, "A", request, claim)
		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "modificarperfil":
			return routers.ModificarPerfil(ctx, claim)

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "eliminartweet":
			return routers.EliminarTweet(request, claim)
		case "bajaRelacion":
			return routers.BajaRelacion(request, claim)
		}
		//
	}

	r.Message = "Method Invalid"
	return r

}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	fmt.Println("funcion  validoAuthorization")
	path := ctx.Value(models.Key("path")).(string)

	if path == "registro" || path == "login" || path == "obteneravatar" || path == "obtenerbanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token Requerido", models.Claim{}
	}

	claim, todoOk, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOk {
		if err != nil {
			fmt.Println("lcano 1")
			fmt.Println("Error en el token" + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("lcano 2")
			fmt.Println("Error en el token" + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim

}
