package service

import (
	"task/Go-000/Week02/model"
	"task/Go-000/Week02/repo"
)

/**
根据ID查找用户
 */
func FindUserById(id int) (*model.UserModel,error){
	return repo.GetUserById(id)
}
