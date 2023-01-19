package main

import (
	"github.com/PedroALeo/crudRedis/database"
	"github.com/PedroALeo/crudRedis/routes"
	"github.com/PedroALeo/crudRedis/scripts"
)

func main() {
	database.ConnectRedis()
	scripts.FlushDB()
	scripts.PopulateDB(nil)
	routes.RequestContollers()
}
