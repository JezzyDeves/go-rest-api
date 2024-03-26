package main

import (
	"github.com/JezzyDeves/go-rest-api/api"
	"github.com/JezzyDeves/go-rest-api/api/routes"
)

func main() {
	routes.InitLoginRoute(api.Echo)
	routes.InitEmployeeRoutes(api.Echo)

	api.Echo.Logger.Fatal(api.Echo.Start(":8080"))
}
