package utl

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/onskycloud/errors"
	"github.com/onskycloud/rbac/model"
	"strings"
)

// ServiceName Seperate Char
const ServiceName = "rbac"

// ParseToken parses token from Authorization header
func ParseToken(key string, algo string, token string) (*model.AuthUser, error) {

	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.BadRequest(ServiceName, "rbac:utl:ParseToken:invalid")
	}
	j,err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(algo) != token.Method {
			return nil, errors.Forbidden(ServiceName, "rbac:utl:ParseToken:forbidden")
		}
		return key, nil
	})
	if err!=nil{
		return nil,err
	}
	claims := j.Claims.(jwt.MapClaims)
	//check cache
	return &model.AuthUser{
		ID:             claims["id"].(int),
		Username:       claims["u"].(string),
		CustomerNumber: claims["c_n"].(string),
		CustomerID:     claims["c"].(int),
		UserUuid:       claims["u_uid"].(string),
		Role:           model.AccessRole(int8(claims["r"].(float64))),
	}, nil
}
