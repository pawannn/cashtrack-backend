package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/ports"
)

type JWTService struct {
	secretKey string
}

func InitJWTService(cfg *config.CashTrackCfg) ports.AuthRepo {
	return JWTService{
		secretKey: cfg.AuthTokenSecret,
	}
}

func (aS JWTService) GenerateUserToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(aS.secretKey)
}

func (aS JWTService) ParseUserToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return aS.secretKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if userID, ok := claims["userID"].(string); ok {
			return userID, nil
		}
		return "", nil
	}
	return "", nil
}
