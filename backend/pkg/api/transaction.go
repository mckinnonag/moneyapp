package api

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// TransactionService contains the methods of the transaction service
type TransactionService interface {
	New(c *gin.Context, tx NewTransactionRequest) error
	Get(c *gin.Context) error
}

// TransactionRepository is what lets our service do db operations without knowing anything about the implementation
type TransactionRepository interface {
	CreateTransaction(NewTransactionRequest) error
}

type transactionService struct {
	storage TransactionRepository
}

func NewTransactionService(transactionRepo TransactionRepository) TransactionService {
	return &transactionService{
		storage: transactionRepo,
	}
}

func (t *transactionService) New(c *gin.Context, tx NewTransactionRequest) error {
	if tx.ID == "" {
		return errors.New("tx service - transaction id required")
	}

	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("request context does not contain user id claim")
	}
	tx.UID = uid.(string)

	err := t.storage.CreateTransaction(tx)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionService) Get(c *gin.Context) ([]NewTransactionRequest, error) {
	uid, exists := c.Get("uid")
	if !exists {
		return nil, errors.New("request context does not contain user id claim")
	}

	tx, err := t.storage.GetTransactions(uid)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
