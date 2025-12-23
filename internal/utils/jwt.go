package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jagoanbunda/jagoanbunda-backend/internal/domain"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID       `json:"user_id"`
	Email  string          `json:"email"`
	Role   domain.UserRole `json:"role"`
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"user_id"`
}

func GenerateAccessToken(userID uuid.UUID, email string, role domain.UserRole) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &AccessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "Jagoan Bunda Backend",
		},
		UserID: userID,
		Email:  email,
		Role:   role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if tokenString, err := token.SignedString([]byte(GetEnv("JWT_SECRET_KEY", "jagoanbundah"))); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func GenerateRefreshToken(userID uuid.UUID) (string, error) {

	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "Jagoan Bunda Backend",
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString([]byte(GetEnv("JWT_SECRET_KEY", "rahasia-dapur kreanova"))); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func ValidateRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(`unexpected signing method : %v`, token.Header["alg"])
		}
		return []byte(GetEnv("JWT_SECRET_KEY", "jagoanbundah")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshTokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid refresh token")
}
