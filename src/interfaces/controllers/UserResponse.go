package controllers

import "github.com/wakabaseisei/runapp/domain"

// クライアントに返すユーザー情報の型
type UserResponse struct {
	Id           int    `json:"id"`
	Sex          int    `json:"sex"`
	Introduction string `json:"introduction"`
}

// usecaseから受け取ったUser型を、クライアントに返す用の型に変換する
func ConvertToUserResponse(user domain.User) UserResponse {
	res := UserResponse{}
	res.Id = user.Id
	res.Sex = user.Sex
	res.Introduction = user.Introduction
	return res
}
