package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword:  es la rutina que permite encriptar password */
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err != nil {
		return err.Error(), err
	}
	return string(bytes), nil
}
