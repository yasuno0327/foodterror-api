package serializer

import "sandbox-api/model"

type SignInResponse struct {
	Token string `json:"api_token"`
}

type UserList struct {
	Users []model.User `json:"users"`
}
