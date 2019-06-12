package authentication

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JwtSecret = []byte("F2AD37825335F43B3170B2C52A4F75C902D58581D2D8F31479F9D13F2253F794")
)

const (
	authenticationHeaderName = "Authorization"
	ExpirationKey            = "exp"
	UserIDKey                = "uid"
	IssuerKey                = "iss"
	JWTIDKey                 = "jti"
	AudienceKey              = "aud"
	SessionDuration          = 8760 * time.Hour // 1 year
)

func ParseSessionJWT(ctx context.Context) (uint64, error) {
	tokenString, _ := extractAuthorizationHeaderFromContext(ctx)
	claims, err := parseJWT(tokenString)
	if err != nil {
		return 0, err
	}

	// issuer := claims[IssuerKey].(string)
	// if issuer != os.Getenv("FRONTEND_ADDRESS") {
	// 	return 0, errors.New("Incorrect issuer")
	// }

	expiration := int64(claims[ExpirationKey].(float64))
	if expiration < time.Now().Unix() {
		return 0, errors.New("Authorization token expired")
	}

	fmt.Println(claims)

	userID := uint64(claims[UserIDKey].(float64))
	return userID, nil
}

// extractAuthorizationHeaderFromContext finds and extracts the Authorization JWT from a context.
func extractAuthorizationHeaderFromContext(ctx context.Context) (string, error) {
	jwt, err := extractJWTFromContext(ctx)
	jwt = strings.Split(jwt, "Bearer ")[1]

	if err != nil {
		return "", err
	}

	return jwt, nil
}

func extractJWTFromContext(ctx context.Context) (string, error) {
	errMissingAuthorizationHeader := fmt.Errorf("missing %q header", authenticationHeaderName)
	t, err := GetAuthFromContext(ctx)
	if len(t) == 0 || err != nil {
		return "", errMissingAuthorizationHeader
	}
	return t, nil
}

func parser(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("Unexpected signing method")
	}
	return JwtSecret, nil
}

func parseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, parser)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Unable to parse JWT")
}
