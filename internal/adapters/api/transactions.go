package api

import (
	"encoding/json"
	"github.com/google/uuid"
	db "main.go/internal/adapters/db"
	"main.go/internal/models"
	"net/http"
	"time"
)

func GetTransactionsCount() int {
	db, ctx := db.ConnectDb()
	var transactions = []models.Transaction{}
	count, err := db.NewSelect().Model(&transactions).Limit(20).ScanAndCount(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

func GetAllTransactions() []models.Transaction {
	db, ctx := db.ConnectDb()
	var transactions = []models.Transaction{}
	err := db.NewSelect().Model(&transactions).Limit(20).Scan(ctx)
	if err != nil {
		panic(err)
	}
	return transactions
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	db, ctx := db.ConnectDb()
	var trans *models.Transaction
	err := json.NewDecoder(r.Body).Decode(&trans)
	trans.ID = uuid.New()
	trans.CreatedAt = time.Now().Format("02-Jan-2006 15:04:05")
	_, err = db.NewInsert().Model(trans).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
