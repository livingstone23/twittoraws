package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/livingstone23/twittoraws/bd"
	"github.com/livingstone23/twittoraws/models"
)

// Estructura necesaria para subir archivo
type readSeeker struct {
	io.Reader
}

// Funcion para subir via lanmbda
func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	fmt.Println("Ingresando en funcion UploadImage")
	var r models.RespApi
	r.Status = 400

	IDUsuario := claim.ID.Hex()

	var filename string
	var usuario models.Usuario

	fmt.Println("funcion_UploadImage antes Bucket name1 :")

	bucket := aws.String(ctx.Value(models.Key("bucketname")).(string))

	fmt.Println("funcion_UploadImage Bucket name: " + strings.ToLower(*bucket))

	switch uploadType {
	case "A":
		filename = "avatars/" + IDUsuario + ".jpg"
		usuario.Avatar = filename
	case "B":
		filename = "banners/" + IDUsuario + ".jpg"
		usuario.Banner = filename
	}

	fmt.Println("funcion_UploadImage Pasando el nombre archivo imagen : " + filename)

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		r.Status = 500
		r.Message = "Error Contentent Type " + err.Error()
		return r
	}

	fmt.Println("funcion_UploadImage Pasando el Content-Type : " + filename)

	if strings.HasPrefix(mediaType, "multipart/") {

		fmt.Println("funcion_UploadImage ingreso if del multipart")

		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			r.Status = 500
			r.Message = "funcion_UploadImage ingreso if del multipart 1" + err.Error() + "Hasta aqui error" + request.Body
			return r
		}

		fmt.Println("funcion_UploadImage Pasando 'multipart'")

		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		p, err := mr.NextPart()
		if err != nil && err != io.EOF {
			r.Status = 500
			r.Message = "funcion_UploadImage ingreso if del multipart 2" + err.Error()
			return r
		}

		fmt.Println("funcion_UploadImage Pasando 'boundary'")

		if err != io.EOF {
			if p.FileName() != "" {
				buf := bytes.NewBuffer(nil)
				if _, err := io.Copy(buf, p); err != nil {
					r.Status = 500
					r.Message = "funcion_UploadImage ingreso if del multipart 3" + err.Error()
					return r
				}

				sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
				if err != nil {
					r.Status = 500
					r.Message = "funcion_UploadImage ingreso if del multipart 4" + err.Error()
					return r
				}

				fmt.Println("funcion_UploadImage Pasando 'us-east-1'")

				uploader := s3manager.NewUploader(sess)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(filename),
					Body:   &readSeeker{buf},
				})

				if err != nil {
					r.Status = 500
					r.Message = "funcion_UploadImage ingreso if del multipart 5" + err.Error()
					return r
				}
			}
		}

		status, err := bd.ModificoRegistro(usuario, IDUsuario)
		if err != nil || !status {
			r.Status = 400
			r.Message = "funcion_UploadImage ingreso if del multipart 6" + err.Error()
			return r
		}

		fmt.Println("funcion_UploadImage Pasando 'ModificoRegistro'")

	} else {
		r.Message = "Debe enviar una imagen con el 'Content-Type' de tipo 'multipart/' en el Header"
		return r
	}

	r.Status = 200
	r.Message = "Imagen Upload OK! "
	return r

}
