package main

import routers "github.com/dictuantranit/go-ecommerce-backend-api/internal/router"

func main() {
	r := routers.NewRouter()

	r.Run(":8002")

	//initialize.Run()
}
