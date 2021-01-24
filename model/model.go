package model

import (
	"SC/token"
	"errors"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// //JwtClaims token生成和验证
// type JwtClaims struct {
// 	jwt.StandardClaims
// 	UserID string `json:"user_id"`
// }

const (
	ErrorReasonServerBusy = "服务器繁忙"
	ErrorReasonReLogin    = "请重新登陆"
)

var Secret = "vinegar" //加醋
//Jwt 在token里有相同的
type Jwt struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func VerifyToken(strToken string) (string, error) {
	//解析
	token, err := jwt.ParseWithClaims(strToken, &token.Jwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return "", errors.New(ErrorReasonServerBusy + ",或token解析失败")
	}
	claims, ok := token.Claims.(*Jwt)
	if !ok {
		return "", errors.New(ErrorReasonReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return "", errors.New(ErrorReasonReLogin)
	}

	return claims.ID, nil //token.Jwt 结构体类型
}
func GetName(id string) string {

	return ""
}
func GetPicture(id string) string {
	return ""
}

//
//func UpdateUserInfo(user User) error {
//	return nil
//}
