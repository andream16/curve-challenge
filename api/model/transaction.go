package model

import (
	"time"

	"github.com/andream16/curve-challenge/pkg/uuid"
	internalError "github.com/andream16/curve-challenge/internal/error"
)

const TimeFormat = "2006-01-02 15:04:05"

// Transaction embeds all transaction information
type Transaction struct {
	// ID is the transaction unique identifier
	ID string `json:"ID"`
	// Receiver is receiver's unique identifier
	Receiver string `json:"receiver"`
	// Sender is receiver's unique identifier
	Sender string `json:"sender"`
	// Amount is the transaction amount
	Amount float64 `json:"amount"`
	// Date is the TS of the action
	Date string `json:"date"`
	// Type is the transaction type
	Type string `json:"type"`
}

// SetReceiver sets transaction's receiver
func (t *Transaction) SetReceiver(receiver string) *Transaction {
	t.Receiver = receiver
	return t
}

// SetSender sets transaction's sender
func (t *Transaction) SetSender(sender string) *Transaction {
	t.Sender = sender
	return t
}

// SetType sets transaction's type
func (t *Transaction) SetType(txType string) *Transaction {
	t.Type = txType
	return t
}

// SetAmount sets transaction's amount
func (t *Transaction) SetAmount(amount float64) *Transaction {
	t.Amount = amount
	return t
}

// SetID sets transaction's ID
func (t *Transaction) SetID() *Transaction {
	t.ID = uuid.New()
	return t
}

// SetDate sets transaction's date
func (t *Transaction) SetDate() *Transaction {
	now := time.Now()
	t.Date = now.Format(TimeFormat)
	return t
}

// NewTransaction creates a new transaction
func NewTransaction(receiver, sender, txType string, amount float64) (*Transaction, error) {

	if len(receiver) == 0 {
		return nil, internalError.Format(Receiver, ErrEmptyParameter)
	}
	if len(sender) == 0 {
		return nil, internalError.Format(Receiver, ErrEmptyParameter)
	}
	if len(txType) == 0 {
		return nil, internalError.Format(TransactionType, ErrEmptyParameter)
	}
	if amount == 0.0 {
		return nil, internalError.Format(Receiver, ErrEmptyParameter)
	}

	out := new(Transaction)
	out.SetID().SetDate().SetAmount(amount).SetReceiver(receiver).SetSender(sender)

	return out, nil

}