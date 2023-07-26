package jwt

import (
	"errors"
	"gowebdemo/configs/appone"
	"time"

	"dario.cat/mergo"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

func Token(kvs jwt.MapClaims) (string, error) {
	jwtConf := appone.GetJwt()
	// exp := time.Now().Add(time.Hour * 24).Unix()
	exp := time.Now().Add(time.Duration(jwtConf.Exp)).Unix()

	claims := jwt.MapClaims{
		"exp": exp,
	}

	if err := mergo.Merge(&claims, kvs); err != nil {
		return "", errors.Join(err, errors.New("merge claims err"))
	}
	log.Info().Msgf("claims:%v", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte("fuckuman"))
	if err != nil {
		log.Error().Msgf("sign token err:%s", err.Error())
		return "", err
	}

	return tokenStr, err
}

func VerifyToken(tokenStr string) (map[string]any, error) {
	parsedToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte("fuckuman"), nil
	})

	if err != nil {
		log.Error().Msgf("解析token失败:%s", err.Error())
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid jwt")
}
