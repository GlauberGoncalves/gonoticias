package migration

import (
	"log"

	mysql "github.com/glaubergoncalves/gonoticias/db"
	"github.com/glaubergoncalves/gonoticias/models"
)

func AutoMigration() bool {
	db, err := mysql.Connect()
	defer db.Close()

	if err != nil {
		log.Println("Não foi possivel fazer migration, banco não disponivel")
		return false
	}
	db.AutoMigrate(models.Noticia{})
	return true
}
