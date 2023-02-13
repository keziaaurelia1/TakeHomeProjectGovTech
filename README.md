# TakeHomeProjectGovTech

## Run The Program

Edit the `.env` first

`go run main.go`

## Generate mock
Mock Repo : `mockgen -package=mock -source=internal/app/$(SUBDOMAIN)/domain/$(SUBDOMAIN)/repo.go -destination=internal/app/$(SUBDOMAIN)/domain/$(SUBDOMAIN)_mock/repository.go`

Mock App : `mockgen -package=mock -source=internal/app/$(SUBDOMAIN)/app/app.go -destination=internal/app/$(SUBDOMAIN)/app_mock/app.go`

## API Documentation 

`Product API.postman_collection.json`

## Notes

Dalam pembuatan project digunakan `wire` yang merupakan Dependency Injection untuk mempermudah pengaturan depedency.

Digunakan `squirrel` untuk mempermudah pembuatan SQL queries
