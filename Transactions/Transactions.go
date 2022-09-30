package Transaction

type CreateTransactionRequest struct {
	Amount   int64  `json:"amount" validate:"notnull,required"`
	Currency string `json:"currency" validate:"notnull,required"`
}
