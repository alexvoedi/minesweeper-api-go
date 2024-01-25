package main

import (
	"os"

	"github.com/alexvoedi/minesweeper-api-go/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	routers.InitRouter().Run(":" + port)
}
