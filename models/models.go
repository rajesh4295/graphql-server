package models

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	WalletId string `json:"walletId"`
}

type Wallet struct {
	Id           string  `json:"id"`
	Balance      int32   `json:"balance"`
	Transactions []int32 `json:"transactions"`
}

type Transaction struct {
	Id        string          `json:"id"`
	Type      TransactionType `json:"type"`
	Amount    int32           `json:"amount"`
	Timestamp string          `json:"timestamp"`
	Status    StatusType      `json:"status"`
}

type TransactionType struct {
	slug string
}

func (t *TransactionType) ToString() string {
	return t.slug
}

func TransactionTypeFromString(transaction string) *TransactionType {
	switch transaction {
	case "unknown":
		return UnknownTransaction
	case "debit":
		return DebitTransaction
	case "credit":
		return CreditTransaction
	default:
		return UnknownTransaction
	}
}

var (
	UnknownTransaction = &TransactionType{"unknown"}
	DebitTransaction   = &TransactionType{"debit"}
	CreditTransaction  = &TransactionType{"credit"}
)

type StatusType struct {
	slug string
}

var (
	UnknownStatus = &StatusType{"unknown"}
	SuccessStatus = &StatusType{"success"}
	FailureStatus = &StatusType{"failure"}
)

func (s *StatusType) ToString() string {
	return s.slug
}

func StatusTypeFromString(status string) *StatusType {
	switch status {
	case "unknown":
		return UnknownStatus
	case "success":
		return SuccessStatus
	case "failure":
		return FailureStatus
	default:
		return UnknownStatus
	}
}
