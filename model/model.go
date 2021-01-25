package model

import (
	"SC/token"
	"errors"

	"github.com/dgrijalva/jwt-go"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
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

//var Secret = "vinegar" //加醋

//Jwt 在token里有相同的
type Jwt struct {
	StudentID string `json:"student_id"`

	jwt.StandardClaims
}

func VerifyToken(strToken string) (string, error) {
	//解析
	token, err := jwt.ParseWithClaims(strToken, &token.Jwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("vinegar"), nil
	})

	if err != nil {
		return "", errors.New(ErrorReasonServerBusy + ",或token解析失败")
	}
	claims, ok := token.Claims.(*Jwt)
	if !ok {
		return "", errors.New(ErrorReasonReLogin + "666")
	}
	if err := token.Claims.Valid(); err != nil {
		return "", errors.New(ErrorReasonReLogin)
	}

	return claims.StudentID, nil //token.Jwt 结构体类型
}

func GetUserInfo(id string) User {
	var u User
	DB.Where("student_id = ?", id).First(&u)
	return u
}

func UpdateUserInfo(user User) error {
	var u User
	DB.Model(&u).Update(user)
	return nil
}

func GetAchievement(id string) []string {
	var achievements []Achievement
	DB.Where("student_id = ?", id).Find(&achievements)
	var achs []string
	for i := 0; i < len(achievements); i++ {
		achs[i] = achievements[i].Achievement
	}
	return achs
}

//Punchs是数组结构体 punch2是[]Punch Punch是为Title与Number的结构体类型
func GetPunchAndNumber(id string) Punchs {
	var punchs []UsersPunch
	DB.Where("student_id = ?", id).Find(&punchs)
	var punchs2 []Punch
	for i := 0; i < len(punchs); i++ {
		punchs2[i].Title = punchs[i].Title
		punchs2[i].Number = punchs[i].Number
	}
	Punchs := Punchs{Punchs: punchs2}
	return Punchs

}
