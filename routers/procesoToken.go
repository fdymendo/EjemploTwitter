package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fdymendo/EjemploTwitter/bd"
	"github.com/fdymendo/EjemploTwitter/models"
)

var Email string
var IDUsuario string

func ProcesoToken(token string) (*models.Claim, bool, string, error) {
	miClave := []byte("ClavePrivada")
	claims := &models.Claim{}

	splitToken := strings.Split(token, " ")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}
	return claims, false, "", err

}
