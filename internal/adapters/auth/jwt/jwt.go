package auth

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/ports"
	"github.com/pawannn/cashtrack/internal/utils"
)

type JWTService struct {
	secretKey string
}

func InitJWTService(cfg *config.CashTrackCfg) ports.AuthRepo {
	return JWTService{
		secretKey: cfg.AuthTokenSecret,
	}
}

func (aS JWTService) GenerateUserToken(userID string) (string, utils.CashTrackError) {
	claims := jwt.MapClaims{
		"userID": userID,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(aS.secretKey)
	if err != nil {
		return "", utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to generate token",
			Error:   err,
		}
	}
	return token, utils.NoErr
}

func (aS JWTService) ParseUserToken(token string) (string, utils.CashTrackError) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return aS.secretKey, nil
	})
	if err != nil {
		return "", utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to parse token",
			Error:   err,
		}
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if userID, ok := claims["userID"].(string); ok {
			return userID, utils.NoErr
		}
		return "", utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to parse token",
			Error:   err,
		}
	}
	return "", utils.CashTrackError{
		Code:    http.StatusInternalServerError,
		Message: "Unable to parse token",
		Error:   err,
	}
}
