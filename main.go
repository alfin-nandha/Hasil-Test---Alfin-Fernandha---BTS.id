package main

import (
	"btsid/test/config"
	"btsid/test/factory"
	"btsid/test/migration"
	"btsid/test/routes"
)

func main() {
	db := config.InitDB()
	migration.Migrate(db)
	presenter := factory.InitFactory(db)
	e := routes.New(presenter)
	e.Logger.Fatal(e.Start(":8000"))
}
