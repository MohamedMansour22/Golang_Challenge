package Transaction

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"net/http"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`
	ID            uuid.UUID        `bun:"id,notnull,pk,type:uuid,default:gen_random_uuid()"`
	Amount        float64          `bun:"amount" `
	Currency      string           `bun:"currency" `
	CreatedAt     pgtype.Timestamp `bun:"createdat" `
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

//func GetTransactions() []Transaction {
//	return Transactions
//}

//var Transactions = []Transaction{
//	{ID: uuid.New(), Amount: 154, Currency: "USD", CreatedAt: time.Now().UTC()},
//	{ID: uuid.New(), Amount: 2400, Currency: "CLP", CreatedAt: time.Now().UTC()},
//	{ID: uuid.New(), Amount: 36400, Currency: "COP", CreatedAt: time.Now().UTC()},
//	{ID: uuid.New(), Amount: 458, Currency: "PEN", CreatedAt: time.Now().UTC()},
//	{ID: uuid.New(), Amount: 3640, Currency: "COP", CreatedAt: time.Now().UTC()},
//}

//func GetAllTransactions() []Transaction {
//	return Transactions
//}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var trans *Transaction
	json.NewDecoder(r.Body).Decode(&trans)
	trans.ID = uuid.New()
	//trans.CreatedAt = time.Now().UTC()
	//if trans.Amount > 0 && trans.Amount <= 100000 {
}
