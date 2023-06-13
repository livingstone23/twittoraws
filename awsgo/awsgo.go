package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

// InicializoAWS permite inicializar 2 variables y arrancar configuracion con amazon
func InicializoAWS() {
	Ctx = context.TODO() //Crea un context vacio
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic("Error al cargar la configuracion .aws/config" + err.Error())
	}

}
