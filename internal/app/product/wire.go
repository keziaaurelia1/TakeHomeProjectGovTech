//go:build wireinject
// +build wireinject

package product

import (
	"github.com/google/wire"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/infra"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/app"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/port"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/repo"
)

func InitHTTP() port.HTTP {
	wire.Build(
		port.ProvideHTTP,
		app.ProvideApplication,
		infra.ProvideMySQL,
		repo.ProvideMySQLRepository,
	)

	return port.HTTP{}
}
