package jwt

import (
	"gvb/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Sign signs the payload with the specified secret.
// The token content.
// iss: （Issuer）签发者
// iat: （Issued At）签发时间，用Unix时间戳表示
// exp: （Expiration Time）过期时间，用Unix时间戳表示
// aud: （Audience）接收该JWT的一方
// sub: （Subject）该JWT的主题
// nbf: （Not Before）不要早于这个时间
// jti: （JWT ID）用于标识JWT的唯一ID
func Sign(playload map[string]any) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["iss"] = global.Conf.Jwt.Issuer
	claims["iat"] = now
	claims["exp"] = now + global.Conf.Jwt.Timeout
	claims["nbf"] = now

	//将自定义需要传入的参数添加到claims中
	for k, v := range playload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Conf.Jwt.Secret))
}
