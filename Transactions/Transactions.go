package Transaction

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"net/http"
	"time"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`
	ID            uuid.UUID `bun:"id,notnull,pk,type:uuid,default:gen_random_uuid()"`
	Amount        float64   `bun:"amount" `
	Currency      string    `bun:"currency" `
	CreatedAt     string    `bun:"createdat" `
}

type CreateTransactionRequest struct {
	Amount   int64  `json:"amount" validate:"notnull,required"`
	Currency string `json:"currency" validate:"notnull,required"`
}

func ConnectDb() (db *bun.DB, ctx context.Context) {
	dsn := "postgresql://mohamedmansour:8v7TLmRapRIxLlwl8Tct6Q@free-tier13.aws-eu-central-1.cockroachlabs.cloud:26257/runny-rhino-2736.tribal?sslmode=verify-full"
	ctx = context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	defer conn.Close(context.Background())
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())

	return db, ctx
}

func GetTransactionsCount() int {
	db, ctx := ConnectDb()
	var transactions = []Transaction{}
	count, err := db.NewSelect().Model(&transactions).Limit(20).ScanAndCount(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

func GetAllTransactions() []Transaction {
	db, ctx := ConnectDb()
	var transactions = []Transaction{}
	err := db.NewSelect().Model(&transactions).Limit(20).Scan(ctx)
	if err != nil {
		panic(err)
	}
	return transactions
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	db, ctx := ConnectDb()
	var trans *Transaction
	//var transactionsmodel *CreateTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&trans)
	//dto.Map(trans, transactionsmodel)
	trans.ID = uuid.New()
	trans.CreatedAt = time.Now().Format("02-Jan-2006 15:04:05")
	_, err = db.NewInsert().Model(trans).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
