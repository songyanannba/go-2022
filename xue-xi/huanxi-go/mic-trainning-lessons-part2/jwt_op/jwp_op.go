package jwt_op

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"mic-training-lessons-part2/internal"
	"mic-training-lessons-part2/log"
	"time"
)

const (
	TokenExpired     = "Token已过期"
	TokenNotValidYet = "Token不再有效"
	TokenMalformed   = "Token非法"
	TokenInvalid     = "Tonken无效"
)

type CustomClaims struct {
	jwt.StandardClaims
	ID          uint32
	NickName    string
	AuthorityIs int32
}

type JWT struct {
	SigninKey []byte
}

func NewJWT() *JWT {
	return &JWT{SigninKey: []byte(internal.AppConf.JWTConfig.SingingKey)}
}

func (j *JWT) GenerateJWT(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(j.SigninKey)
	if err != nil {
		log.Logger.Error("生成JWT错误:" + err.Error())
		return "", err
	}
	return tokenStr, nil
}

func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigninKey, nil
	})
	if err != nil {
		if result, ok := err.(jwt.ValidationError); ok {
			if result.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(TokenMalformed)
			} else if result.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(TokenExpired)
			} else if result.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(TokenNotValidYet)
			} else {
				return nil, errors.New(TokenInvalid)
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New(TokenInvalid)
	}
	return nil, errors.New(TokenInvalid)
}

func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigninKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(7 * 24 * time.Hour).Unix()
		return j.GenerateJWT(*claims)
	}
	return "", errors.New(TokenInvalid)
}
