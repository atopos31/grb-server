package jwt

import (
	"errors"
	"gvb/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
)

// Payload is the data of the JSON web token.
type Payload struct {
	Username string
	Password string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the payloads if the token was valid.
func Parse(tokenString string, secret string) (*Payload, error) {
	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return nil, err
	}
	// Read the token if it's valid.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloads := &Payload{}
		payloads.Username = claims["username"].(string)
		payloads.Password = claims["password"].(string)
		return payloads, nil
	}

	// Other errors.
	return nil, err
}

// Sign signs the payload with the specified secret.
// The token content.
// iss: （Issuer）签发者
// iat: （Issued At）签发时间，用Unix时间戳表示
// exp: （Expiration Time）过期时间，用Unix时间戳表示
// aud: （Audience）接收该JWT的一方
// sub: （Subject）该JWT的主题
// nbf: （Not Before）不要早于这个时间
// jti: （JWT ID）用于标识JWT的唯一ID
func Sign(playload Payload) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["iss"] = global.Conf.Jwt.Issuer
	claims["iat"] = now
	claims["exp"] = now + global.Conf.Jwt.Timeout
	claims["nbf"] = now

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Conf.Jwt.Secret))
}
