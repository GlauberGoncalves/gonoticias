package controllers

import (
	"github.com/glaubergoncalves/gonoticias/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

//Index : index
func Home(w http.ResponseWriter, r *http.Request) {
	noticia := models.Noticia{}
	noticias, err := noticia.BuscaTodasAsNoticias(); if err != nil {
		log.Fatal(err)
	}
	temp.ExecuteTemplate(w, "Index", noticias)
}

//New : new
func NovaNoticia(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		noticia := models.Noticia{}
		noticia.Titulo = r.FormValue("titulo")
		noticia.Conteudo = r.FormValue("conteudo")
		noticia.CreatedAt = time.Now()
		noticia.UpdatedAt = time.Now()

		_ , err := noticia.AdicionaNoticia()
		checaErro(err)

		http.Redirect(w, r, "/", 301)
	}
}

//Formulario : Formulario
func Formulario(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Formulario", nil)
}

func checaErro(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
