package script

import (
	"SC/model"
	"encoding/base64"
	"fmt"
)

func Base2Hash() {
	users := getAllUser()
	fmt.Println(len(users))
	for _, user := range users {
		password, _ := base64.StdEncoding.DecodeString(user.Password)
		hashPwd := model.GeneratePasswordHash(string(password))
		model.DB.Table("users").Where("student_id = ?", user.StudentID).Update("password", hashPwd)
	}
}

func getAllUser() []model.User {
	var users []model.User
	model.DB.Find(&users)
	return users
}
