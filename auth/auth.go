package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gitlab.com/voltunes/api-master-project/exception"
	"gitlab.com/voltunes/api-master-project/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CreateAuthFunc func(userID string, tokenDetails *TokenDetails)

type AccessDetails struct {
	UserID    string
	ID        uint
	Role      string
	Status    string
	UnitID    string
	CompanyID uint
	UnitIDs   []string
}

type TokenDetails struct {
	AccessToken        string
	RefreshToken       string
	AccessUUID         string
	RefreshUUID        string
	UserAgent          string
	RemoteAddress      string
	AtExpired          int64
	RefeshTokenExpired int64
}

func Auth(next func(c *gin.Context, auth *AccessDetails), roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check JWT Token
		tokenAuth, err := ExtractTokenMetadata(c.Request)
		if err != nil {
			helper.PanicIfError(exception.ErrUnauthorized)
		}

		// Check Permission User
		// if !helper.Contains(roles, tokenAuth.Role) {
		// 	helper.PanicIfError(exception.ErrPermissionDenied)
		// }

		next(c, tokenAuth)
	}
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		unitIDsInterface, _ := claims["unit_ids"].([]interface{})
		unitIDs := make([]string, len(unitIDsInterface))
		for i, v := range unitIDsInterface {
			unitIDs[i] = v.(string)
		}
		return &AccessDetails{
			UserID:    claims["id"].(string),
			ID:        uint(claims["user_id"].(float64)),
			Role:      claims["role"].(string),
			Status:    claims["status"].(string),
			UnitID:    claims["unit_id"].(string),
			CompanyID: uint(claims["company_id"].(float64)),
			UnitIDs:   unitIDs,
		}, nil
	}
	return nil, err
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
