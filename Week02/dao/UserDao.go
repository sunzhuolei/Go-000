package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"strconv"
	"task/Go-000/Week02/model"
	"task/Go-000/Week02/tool"
)

type UserDao struct{

}

func NewUserDao() *UserDao{
	return &UserDao{}
}


func (user *UserDao)GetUserById(id int)(userModel *model.UserModel,err error){
	cfg,err := tool.ParseConfig("./Go-000/Week02/config/app.json")
	if err != nil{
		return
	}
	database := cfg.Database
	dataSourceName := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" +database.DbName + "?charset=" + database.Charset
	db,err := sql.Open(database.Driver,dataSourceName)
	if err != nil{
		return nil,err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		return
	}
	var (
		uid int
		name string
		phone string
		address string
		sex int
	)
	rows := db.QueryRow("select * from user where id = "+strconv.Itoa(id))
	err = rows.Scan(&uid,&name,&phone,&address,&sex)
	if err == sql.ErrNoRows{
		err = nil
		return
	}
	if err != nil{
		err = errors.Wrap(err,"查询数据失败！")
	}
	userModel = &model.UserModel{uid,name,phone,address,sex}
	return userModel,err
}
