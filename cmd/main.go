package main

import (
	"encoding/json"
	"main.go/internal/adapters/api"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/TransactionsCount", func(w http.ResponseWriter, r *http.Request) {
		transactions := api.GetTransactionsCount()
		json.NewEncoder(w).Encode(transactions)
	})
	r.Get("/Transactions", func(w http.ResponseWriter, r *http.Request) {
		transactions := api.GetAllTransactions()
		json.NewEncoder(w).Encode(transactions)
	})

	r.Post("/Transactions/create", api.CreateTransaction)

	http.ListenAndServe(":3000", r)
}
