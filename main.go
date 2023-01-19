package main

import "github.com/PedroALeo/crudRedis/database"

func main() {
	database.ConnectRedis()
	database.FlushDB()
	database.PopulateDB(nil)
}
