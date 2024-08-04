package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/entities"
)

func DecodeIDToken(idTokenJWT string, role string, status string) (*entities.SaveUserData, error) {

	parts := strings.Split(idTokenJWT, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token parts")
	}

	claimsPart := parts[1]

	claimsJSON, err := base64.RawURLEncoding.DecodeString(claimsPart)
	if err != nil {
		return nil, errors.New("invalid token claims")
	}

	var claims map[string]interface{}
	err = json.Unmarshal(claimsJSON, &claims)
	if err != nil {
		return nil, errors.New("invalid token unmarshal")
	}

	fullName, nameOK := claims["name"].(string)
	email, emailOK := claims["email"].(string)
	sub, subOK := claims["sub"].(string)
	if !nameOK || !emailOK || !subOK {
		return nil, errors.New("invalid token claims")
	}
	var userData entities.SaveUserData = entities.SaveUserData{FullName: fullName, Email: email,
		Preferences: &entities.Preferences{}, Role: role, Status: status, Id: sub}

	return &userData, nil
}
