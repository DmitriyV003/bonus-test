package main

import "bonus-test/app"

func main() {
	app.InitLogger()
	err := app.InitDb()
	if err != nil {
		panic(err)
	}

	server := app.NewServer()
	server.InitServer()
}
