package main

import (
	"github.com/gabrielaPezzini/loja-digport-backend/model"
	"github.com/gabrielaPezzini/loja-digport-backend/tree/main/model"
)

func criaEstoque() []model.Product {
	products := []model.Product{
		{
			Name:        "Product 1",
			Description: "Descrição do Produto 1",
			Category:    "Categoria 1",
			Id:          "1",
			Price:       10.99,
			Quantity:    100,
			Image:       "produto1.jpg",
		},
		{
			Name:        "Produto 2",
			Description: "Descrição do Produto 2",
			Category:    "Catogoria 2",
			Id:          "2",
			Price:       20.99,
			Quantity:    50,
			Image:       "produto2.jpg",
		},
	}
	return products
}
