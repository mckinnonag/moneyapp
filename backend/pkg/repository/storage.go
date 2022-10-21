package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"moneyapp/pkg/api"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Storage interface {
	RunMigrations(connectionString string) error
	CreateTransactions(tx []api.Transaction) error
	GetTransactions(uid string) ([]api.Transaction, error)
	CreateAccessToken(request *api.NewAccessTokenRequest) error
	GetAccessTokens(uid string) ([]string, error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

// Migrate the database
func (s *storage) RunMigrations(connectionString string) error {
	if connectionString == "" {
		return errors.New("repository: the connString was empty")
	}

	// get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	migrationsPath := filepath.Join("file://", basePath, "/pkg/repository/migrations/")

	fmt.Println(connectionString)
	m, err := migrate.New(migrationsPath, connectionString)

	if err != nil {
		return err
	}

	err = m.Up()

	if err == migrate.ErrNoChange {
		return nil
	}

	return err
}

// Create a new transaction
func (s *storage) CreateTransactions(txs []api.Transaction) error {
	txn, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("transactions", "plaid_id", "plaid_item_id", "user_id", "category", "location", "tx_name", "amount", "iso_currency_code", "tx_date", "pending", "merchant_name", "payment_channel"))
	if err != nil {
		return err
	}

	for _, t := range txs {
		_, err = stmt.Exec(t.ID, t.ItemID, t.UID, t.Category, t.Location, t.Name, t.Amount, t.IsoCurrencyCode, t.Date, t.Pending, t.MerchantName, t.PaymentChannel)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Get all of a user's transactions
func (s *storage) GetTransactions(uid string) ([]api.Transaction, error) {
	return nil, errors.New("not implemented")
}

// Store a new access token for a user
func (s *storage) CreateAccessToken(request *api.NewAccessTokenRequest) error {
	sqlStatement := `
		INSERT INTO plaid_items (user_id, access_token, plaid_item_id)
		VALUES ($1, $2, $3)`
	_, err := s.db.Exec(sqlStatement, request.UID, request.AccessToken, request.ItemId)
	if err != nil {
		return err
	}
	return nil
}

// Get all access tokens for a user
func (s *storage) GetAccessTokens(uid string) ([]string, error) {
	var accessToken string
	var result []string
	sqlStatement := `SELECT access_token FROM plaid_items WHERE user_id=$1`
	row := s.db.QueryRow(sqlStatement, uid)
	switch err := row.Scan(&accessToken); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("no access tokens found for uid %s", uid)
	case nil:
		result = append(result, accessToken)
	default:
		return nil, err
	}
	return result, nil
}
