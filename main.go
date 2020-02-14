package main

import "github.com/alditiadika/go-graphiql/app"

func main() {
	apps := app.App{}
	apps.Init()
	apps.Run()

}
