package repo

import (
	"task/Go-000/Week02/dao"
	"task/Go-000/Week02/model"
)

/**
根据用户ID查找用户信息
 */
func GetUserById(id int) (*model.UserModel,error){
	return dao.NewUserDao().GetUserById(id)
}
