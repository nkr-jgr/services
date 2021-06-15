package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	db "github.com/micro/services/db/proto"
	"github.com/micro/services/pkg/tracing"
	"github.com/micro/services/user/handler"
	proto "github.com/micro/services/user/proto"
)

func main() {
	service := service.New(
		service.Name("user"),
	)

	service.Init()

	proto.RegisterUserHandler(service.Server(), handler.NewUser(db.NewDbService("db1", service.Client())))
	traceCloser := tracing.SetupOpentracing("user")
	defer traceCloser.Close()

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
