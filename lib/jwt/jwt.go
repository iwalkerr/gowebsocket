package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jwtSecret     = "&WJof0jaY4ByTHR2"
	jwtExpireTime = 5 * time.Minute //2 * time.Hour
)

type Claims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
	RefreshTime int64 //【The refresh time】 该jwt刷新的时间；unix时间戳
}

// 生成token
func GenerateToken(uid int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(jwtExpireTime)
	refreshTime := nowTime.Add(jwtExpireTime / 2)

	claims := Claims{
		Uid:         uid,
		RefreshTime: refreshTime.Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "web_gin_template",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))
	return token, err
}

// 验证token
// 验证token有效期，如果time过期返回false
// token超过刷新时间，返回新的token
func ParseToken(token string) (int, string, bool) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(jwtSecret), nil
	})
	if tokenClaims == nil || err != nil {
		return 0, "", false
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok || !tokenClaims.Valid {
		return 0, "", false
	}

	// 超过刷新时间
	if time.Now().Unix() > claims.RefreshTime {
		if token, err = GenerateToken(claims.Uid); err == nil {
			return claims.Uid, token, true
		}
	}
	return claims.Uid, "", true
}
