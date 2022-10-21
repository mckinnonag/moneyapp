package api

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// TransactionService contains the methods of the transaction service
type TransactionService interface {
	New(c *gin.Context, txs []Transaction) error
	Get(c *gin.Context) ([]Transaction, error)
}

// TransactionRepository is what lets our service do db operations without knowing anything about the implementation
type TransactionRepository interface {
	CreateTransactions([]Transaction) error
	GetTransactions(string) ([]Transaction, error)
}

type transactionService struct {
	storage TransactionRepository
}

func NewTransactionService(transactionRepo TransactionRepository) TransactionService {
	return &transactionService{
		storage: transactionRepo,
	}
}

func (t *transactionService) New(c *gin.Context, txs []Transaction) error {
	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("request context does not contain user id claim")
	}

	for _, t := range txs {
		if t.ID == "" {
			return errors.New("tx service - transaction id required")
		}
		t.UID = uid.(string)
	}

	err := t.storage.CreateTransactions(txs)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionService) Get(c *gin.Context) ([]Transaction, error) {
	uid, exists := c.Get("uid")
	if !exists {
		return nil, errors.New("request context does not contain user id claim")
	}

	tx, err := t.storage.GetTransactions(uid.(string))
	if err != nil {
		return nil, err
	}
	return tx, nil
}
