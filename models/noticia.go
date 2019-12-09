package models

import (
	"encoding/json"
	database "github.com/glaubergoncalves/gonoticias/db"
	"log"
	"time"
)

type Noticia struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Titulo    string    `gorm:"size:255;not null" json:"titulo"`
	Conteudo  string    `gorm:"size:255;not null" json:"conteudo"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (n *Noticia) BuscaTodasAsNoticias() ([]Noticia, error) {

	conexao, _ := database.Connect()
	defer conexao.Close()

	noticias := []Noticia{}
	reply, err := database.Get("noticias")

	if err != nil {
		log.Println("Buscando no mysql")
		var err error
		noticias := []Noticia{}
		err = conexao.Debug().Model(&Noticia{}).Limit(100).Find(&noticias).Error
		if err != nil {
			return []Noticia{}, err
		}
		noticiasBytes, _ := json.Marshal(noticias)
		database.Set("noticias", noticiasBytes)
		return noticias, nil
	} else {
		log.Println("Buscando no redis")
		json.Unmarshal(reply, &noticias)
		return noticias, nil
	}
}

func (n *Noticia) AdicionaNoticia() (*Noticia, error) {

	conexao, _ := database.Connect()
	defer conexao.Close()

	var err error
	err = conexao.Debug().Model(&Noticia{}).Create(&n).Error
	if err != nil {
		return &Noticia{}, err
	}
	database.Flush("noticias")
	return n, nil
}
