// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ZQCard/kratos-crud-layout/internal/biz"
	"github.com/ZQCard/kratos-crud-layout/internal/conf"
	"github.com/ZQCard/kratos-crud-layout/internal/data"
	"github.com/ZQCard/kratos-crud-layout/internal/server"
	"github.com/ZQCard/kratos-crud-layout/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, auth *conf.Auth, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	db := data.NewMysqlCmd(confData, logger)
	cmdable := data.NewRedisCmd(confData, logger)
	dataData, cleanup, err := data.NewData(db, cmdable, logger)
	if err != nil {
		return nil, nil, err
	}
	ServiceNameRepo := data.NewServiceNameRepo(dataData, logger)
	ServiceNameUseCase := biz.NewServiceNameUseCase(ServiceNameRepo, logger)
	ServiceNameService := service.NewServiceNameService(ServiceNameUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, ServiceNameService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
