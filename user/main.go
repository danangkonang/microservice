package main

import (
	"github.com/danangkonang/user/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app.Run()
}
