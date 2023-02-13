package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/app/product/port"
)

func main() {
	RunHTTPServer()
}

// var router *chi.Mux
// var db *sql.DB

// const (
// 	dbName = "go-mysql-crud"
// 	dbPass = "12345"
// 	dbHost = "localhost"
// 	dbPort = "33066"
// )

// func routers() *chi.Mux {
// 	router.Get("/posts", AllPosts)
// 	router.Get("/posts/{id}", DetailPost)
// 	router.Post("/posts", CreatePost)
// 	router.Put("/posts/{id}", UpdatePost)
// 	router.Delete("/posts/{id}", DeletePost)

// 	return router
// }

// func init() {
// 	router = chi.NewRouter()
// 	router.Use(middleware.Recoverer)

// 	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)

// 	var err error
// 	db, err = sql.Open("mysql", dbSource)

// 	catch(err)
// }
// func Logger() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(time.Now(), r.Method, r.URL)
// 		router.ServeHTTP(w, r) // dispatch the request
// 	})
// }

func RunHTTPServer() {
	portUse := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		portUse = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", portUse)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})
	new := product.InitHTTP()
	tmp := port.ProvideHTTP(new.App)
	r.Mount("/products", tmp.RegisterRouter())
	// http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe(":"+portUse, r))
}
