package token_verify

import (
	"fmt"
	"ginchat/global/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(config.Conf.Token.JwtTokenSignKey)

/**
 * @description 生成token
 * @return
 * @date 2023/12/12 14:21:28 @version 1.0.0 @add jz.wang
 */
func GenerateToken(userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Second * config.Conf.Token.JwtTokenCreatedExpireAt)

	claims := CustomClaims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			Issuer:    config.Conf.Token.JwtIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	fmt.Printf("now time %s, expore time %s, token: %s, config: %d", nowTime, expireTime, token, config.Conf.Token.JwtTokenCreatedExpireAt)

	return token, err
}

/**
 * @description 校验token
 * @return
 * @date 2023/12/12 14:37:56 @version 1.0.0 @add
 */
func ParseToken(tokenString string) (*CustomClaims, error) {
	fmt.Println("ParseToken :", tokenString)
	tokenClaims, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if tokenClaims != nil {
		fmt.Println("tokenClaims 1===:", tokenClaims, "===============, ", tokenClaims.Valid)
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			fmt.Println("tokenString 2===:", tokenString)
			return claims, nil
		}
	}

	return nil, err
}
