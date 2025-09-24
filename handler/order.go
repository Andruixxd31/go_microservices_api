package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (order *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Created an order")
}

func (order *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listing orders")
}

func (order *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got an order")
}

func (order *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleted an order")
}

func (order *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updated an order")
}
