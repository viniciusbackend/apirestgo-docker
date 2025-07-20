package router

import (
	"log"
	"net/http"
	"testapi/controllers"
	"testapi/middleware"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/", controllers.Inicio)
	r.HandleFunc("/produtos", controllers.ListarProdutos).Methods("Get")
	r.HandleFunc("/produtos/{id}", controllers.ListarProdutoPorId).Methods("Get")
	r.HandleFunc("/produtos", controllers.CriarProduto).Methods("Post")
	r.HandleFunc("/produtos/{id}", controllers.DeletarProduto).Methods("Delete")
	r.HandleFunc("/produtos/{id}", controllers.EditarProduto).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
