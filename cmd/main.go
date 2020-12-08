package main

import (
	"log"
	"net/http"
	"net/url"

	"solution/infrastructure"
	"solution/routes"

	"github.com/go-chi/chi"
)

func main() {
	route := chi.NewRouter()
	r := routes.RegisterRoutes(route)
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}
