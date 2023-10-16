package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("XuZh-Owner")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "XuZh",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret) //数字签名操作 对于HS256加密算法来说，需要与jwtSecret进行一个映射的操作.
	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	//以及一个回调函数用于提供 JWT 密钥（jwtSecret）
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
		/*
			在这段代码中，我们使用类型断言语法 tokenClaims.Claims.(*Claims) 将 tokenClaims.Claims 断言为 *Claims 类型。如果断言成功，也就是 tokenClaims.
			Claims 是 *Claims 类型，那么变量 claims 将获得断言后的值，并且 ok 的值将为 true。
		*/
	}
	return nil, err
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateEmailToken 签发 Email token
func GenerateEmailToken(userId, Operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := EmailClaims{
		UserID:        userId,
		Email:         email,
		Password:      password,
		OperationType: Operation,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "XuZh",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret) //数字签名操作 对于HS256加密算法来说，需要与jwtSecret进行一个映射的操作.
	return token, err
}

// ParseEmailToken 验证 Email token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	//以及一个回调函数用于提供 JWT 密钥（jwtSecret）
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
		/*
			在这段代码中，我们使用类型断言语法 tokenClaims.Claims.(*Claims) 将 tokenClaims.Claims 断言为 *Claims 类型。如果断言成功，也就是 tokenClaims.
			Claims 是 *Claims 类型，那么变量 claims 将获得断言后的值，并且 ok 的值将为 true。
		*/
	}
	return nil, err
}
