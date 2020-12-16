package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	api_user_v1 "task/Go-000/Week04/api/user/v1"
	"task/Go-000/Week04/internal/server"
)


func main() {

	srv, cleanup, err := server.InitializeServer()
	defer cleanup()
	if err != nil {
		log.Printf("Init Server error:%v\n", err)
		return
	}
	//go clientGrpc()

	log.Println("Start Server")
	if err = srv.Run(); err != nil {
		log.Printf("Run Server error:%v\n", err)
		return
	}
}

/**
客户端调用
 */
func clientGrpc(){
	conn,err := grpc.Dial(":8090",grpc.WithInsecure())
	if err != nil{
		log.Fatal(err)
		return
	}
	defer conn.Close()
	client := api_user_v1.NewUserServiceClient(conn)
	response,err := client.GetUser(context.Background(),
		&api_user_v1.UserRequest{Uid: 1},
	)
	if err != nil{
		log.Fatal(err)
		return
	}
	log.Println(response.Name,response.Age)
}
