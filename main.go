package main

import (
	"github.com/glaubergoncalves/gonoticias/db/migration"
	"github.com/glaubergoncalves/gonoticias/routes"
	"net/http"
)

func main() {

	migration.AutoMigration()
	routes.CarregaRotas()

	http.ListenAndServe(":8080", nil)
}