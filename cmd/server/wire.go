//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/ZQCard/kratos-crud-layout/internal/biz"
	"github.com/ZQCard/kratos-crud-layout/internal/conf"
	"github.com/ZQCard/kratos-crud-layout/internal/data"
	"github.com/ZQCard/kratos-crud-layout/internal/server"
	"github.com/ZQCard/kratos-crud-layout/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
