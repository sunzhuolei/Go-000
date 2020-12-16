package server

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	api_user_v1 "task/Go-000/Week04/api/user/v1"
	"task/Go-000/Week04/internal/service"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)
type Config struct{
	Database DatabaseConig `json:"database"`
	Grpc GrpcConfig
}

type GrpcConfig struct {
	Port string
}

type DatabaseConig struct {
	Driver string `json:"driver"`
	User string `json:"user"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
	DbName string `json:"db_name"`
	Charset string `json:"charset"`
	ShowSql bool `json:"show_sql"`
}
var _Cfg *Config = nil

func GetConfig() *Config{
	return _Cfg
}
func ParseConfig(path string)(*Config,error){

	file,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err :=decoder.Decode(&_Cfg);err != nil{
		return nil,err
	}
	return _Cfg,nil
}
type Server struct {
	service *service.UserService
}

func NewServer(s *service.UserService) *Server {
	return &Server{service: s}
}

/**
启动服务
 */
func (srv *Server) Run() error {
	cfg,err := ParseConfig("./Go-000/Week04/config/app.json")
	if err != nil{
		log.Fatal(err.Error())
		return err
	}
	lis, err := net.Listen("tcp", cfg.Grpc.Port)
	if err != nil {
		return err
	}
	g, ctx := errgroup.WithContext(context.Background())
	s := grpc.NewServer()
	g.Go(func() error {
		go func() {
			<-ctx.Done()
			s.GracefulStop()
			log.Printf("Shutdown Server")
		}()
		api_user_v1.RegisterUserServiceServer(s, srv.service)
		return s.Serve(lis)
	})
	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			select {
			case <-ctx.Done():
				return nil
			case s := <-c:
				log.Printf("get a signal %s", s.String())
				switch s {
				case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
					return errors.New("Close by signal " + s.String())
				case syscall.SIGHUP:
				default:
					return errors.New("Undefined signal")
				}
			}
		}
	})
	return g.Wait()
}
