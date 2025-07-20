package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testapi/database"
	"testapi/models"

	"github.com/gorilla/mux"
)

func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Tela inicial")
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
	var produtos []models.Produto
	database.DB.Find(&produtos)
	json.NewEncoder(w).Encode(produtos)
}

func ListarProdutoPorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produtoBuscado models.Produto
	database.DB.First(&produtoBuscado, id)
	json.NewEncoder(w).Encode(produtoBuscado)
}

func CriarProduto(w http.ResponseWriter, r *http.Request) {
	var novoProduto models.Produto
	json.NewDecoder(r.Body).Decode(&novoProduto)
	database.DB.Create(&novoProduto)
	json.NewEncoder(w).Encode(novoProduto)
}

func EditarProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produto models.Produto
	database.DB.First(&produto, id)
	json.NewDecoder(r.Body).Decode(&produto)
	database.DB.Save(&produto)
	json.NewEncoder(w).Encode(produto)
}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produto models.Produto
	database.DB.Delete(&produto, id)
	json.NewEncoder(w).Encode(produto)
}
