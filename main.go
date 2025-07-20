package main

import (
	"testapi/database"
	"testapi/router"
)

func main() {
	database.ConectaComBancoDeDados()
	router.Router()
}
