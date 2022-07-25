package routes

import (
	"alexandrefreitasp/controllers"
	"net/http"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index)

}
