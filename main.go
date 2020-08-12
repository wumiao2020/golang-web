package main

import (
	"./routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	routes.FrontendStart()
}
