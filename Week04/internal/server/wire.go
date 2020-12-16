//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/google/wire"
	"task/Go-000/Week04/internal/dao"
	"task/Go-000/Week04/internal/service"
)

func InitializeServer() (*Server, func(), error) {
	wire.Build(NewServer, wire.NewSet(service.NewUserService, dao.Provider))
	return nil, nil, nil
}