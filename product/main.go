package main

import (
	"github.com/danangkonang/product/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app.Run()
}
