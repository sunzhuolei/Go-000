package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"task/Go-000/Week02/service"
)

func main() {
	model,err := service.FindUserById(3)
	if err != nil{
		fmt.Printf("%T %+v\n",err)
		return
	}
	fmt.Println(model)
}
