package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	AUTHORIZATION = "Authorization"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func getPemCert(token *jwt.Token, audience string) (string, error) {
	cert := ""
	resp, err := http.Get(fmt.Sprintf(`%s/.well-known/jwks.json`, audience))
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := fmt.Errorf("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

func JWTAuth0(audience string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get(AUTHORIZATION)
		if authorizationHeader == "" {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}
		tokenParts := strings.Split(authorizationHeader, "Bearer ")
		if len(tokenParts) != 2 {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}
		tokenString := tokenParts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			cert, err := getPemCert(token, audience)
			if err != nil {
				return nil, err
			}
			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), true) {
			sub := claims["sub"]
			ctx.Request.Header.Add("sub", sub.(string))
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
		}
	}
}
