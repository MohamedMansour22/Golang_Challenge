package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"log"
	DB "main.go/internal/adapters/db"
	"main.go/internal/models"
	"net/http"
	"time"
)

var db, ctx = DB.ConnectDb()

func GetTransactionsCount(w http.ResponseWriter, r *http.Request) {
	var transactions = []models.Transaction{}
	count, err := db.NewSelect().Model(&transactions).Limit(20).ScanAndCount(ctx)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(count)
	w.WriteHeader(200)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions = []models.Transaction{}
	err := db.NewSelect().Model(&transactions).Limit(20).Scan(ctx)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(transactions)
	w.WriteHeader(200)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var trans *models.Transaction
	err := json.NewDecoder(r.Body).Decode(&trans)
	trans.ID = uuid.New()
	trans.CreatedAt = time.Now().Format("02-Jan-2006 15:04:05")
	_, err = db.NewInsert().Model(trans).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func UpdateTransaction(trans *models.Transaction) (bool, error) {
	trans.Status = true
	_, err := db.NewUpdate().Model(trans).Column("status").Exec(ctx)
	if err != nil {
		return true, err
	}
	return false, nil
}

func HandleRequest() {
	Router := chi.NewRouter()
	Router.Use(middleware.Logger)
	Router.Get("/TransactionsCount", GetTransactionsCount)
	Router.Get("/Transactions", GetAllTransactions)
	Router.Post("/Transactions/create", CreateTransaction)
	log.Fatal(http.ListenAndServe(":3000", Router))
}
