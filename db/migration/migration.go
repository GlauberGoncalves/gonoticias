package migration

import (
	mysql "github.com/glaubergoncalves/gonoticias/db"
	"github.com/glaubergoncalves/gonoticias/models"
)

func AutoMigration(){
	db := mysql.Connect()
	defer db.Close()
	db.AutoMigrate(models.Noticia{})
}
