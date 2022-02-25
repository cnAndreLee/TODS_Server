package dto

import "github.com/cnAndreLee/tods_server/model"

type UserDto struct {
	Account string `json:"account"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Account: user.Account,
	}
}
