package routes

import (
	"github.com/glaubergoncalves/gonoticias/controllers"
	"net/http"
)

// CarregaRotas : carrega as rotas da aplicação
func CarregaRotas() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/cria-noticia", controllers.NovaNoticia)
	http.HandleFunc("/formulario-noticia", controllers.Formulario)
}