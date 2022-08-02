package Transaction

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `json:"Id" validate:"required"`
	Amount    float64   `json:"Amount" validate:"required,gte=0,lte=100000"`
	Currency  string    `json:"Currency" validate:"required"`
	CreatedAt time.Time `json:"CreatedAt" validate:"required"`
}

var Transactions = []Transaction{
	{ID: uuid.New(), Amount: 154, Currency: "USD", CreatedAt: time.Now().UTC()},
	{ID: uuid.New(), Amount: 2400, Currency: "CLP", CreatedAt: time.Now().UTC()},
	{ID: uuid.New(), Amount: 36400, Currency: "COP", CreatedAt: time.Now().UTC()},
	{ID: uuid.New(), Amount: 458, Currency: "PEN", CreatedAt: time.Now().UTC()},
	{ID: uuid.New(), Amount: 3640, Currency: "COP", CreatedAt: time.Now().UTC()},
}

func GetAllTransactions() []Transaction {
	return Transactions
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var trans *Transaction
	json.NewDecoder(r.Body).Decode(&trans)
	trans.ID = uuid.New()
	trans.CreatedAt = time.Now().UTC()
	if trans.Amount > 0 && trans.Amount <= 100000 {
		Transactions = append(Transactions, *trans)
	} else {
		w.Write([]byte("The entered amount is not vaild!!"))
	}

}
