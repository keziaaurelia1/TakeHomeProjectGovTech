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

Dalam pembuatan project digunakan `wire` yang merupakan Dependency Injection untuk mempermudah pengaturan depedency. Wire ini digunakan untuk mempermudah mengatasi permasalahan dependency. Sehingga saat pembuatan variabel jika terdapat variabel a yang memiliki dependency dengan variabel b maka variabel b akan dibuat terlebih dahulu yang dibutuhkan oleh a kemudian baru membuat a. Dengan begitu wire ini akan mengatasi permasalahan dependency yang mungkin terjadi.

Digunakan `squirrel` untuk mempermudah pembuatan SQL queries. Squirrel juga memudahkan dalam membaca code sehingga dapat dimengerti dengan lebih cepat. Selain itu dari segi security squirrel ini digunakan untuk menghindari sql injection.
