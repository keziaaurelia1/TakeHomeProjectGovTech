// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package product

import (
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/infra"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/app"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/port"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/repo"
)

// Injectors from wire.go:

func InitHTTP() port.HTTP {
	db := infra.ProvideMySQL()
	repository := repo.ProvideMySQLRepository(db)
	application := app.ProvideApplication(repository)
	http := port.ProvideHTTP(application)
	return http
}