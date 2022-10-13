package repository

import (
	"database/sql"
	"errors"

	"moneyapp/pkg/api"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Storage interface {
	RunMigrations(connectionString string) error
	CreateTransaction(request api.NewTransactionRequest) error
	GetTransactions(uid string) ([]api.NewTransactionRequest, error)
	CreateAccessToken(request api.NewAccessTokenRequest) error
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

	m, err := migrate.New(migrationsPath, connectionString)

	if err != nil {
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil
	}

	return nil
}

// Helper function. Accepts 3 char currency code; returns ID in database
func (s *storage) getCurrencyCode(code string) (currency_id string, err error) {
	sqlStatement := `SELECT currency_id, code FROM users WHERE code=$1;`
	row := s.db.QueryRow(sqlStatement, code)
	switch err = row.Scan(&currency_id, &code); err {
	case sql.ErrNoRows:
		return "", errors.New("code does not exist")
	case nil:
		return currency_id, nil
	default:
		return "", err
	}
}

// Create a new transaction
func (s *storage) CreateTransaction(request api.NewTransactionRequest) error {
	iso_currency_code, err := s.getCurrencyCode(request.IsoCurrencyCode)
	if err != nil {
		return err
	}
	sqlStatement := `
		INSERT INTO transactions (plaid_id, plaid_item_id, user_id, category, location, tx_name, amount, iso_currency_code, tx_date, pending, merchant_name, payment_channel)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		`
	_, err = s.db.Exec(sqlStatement, request.ID, request.ItemID, request.UID, request.Category, request.Location, request.Name, request.Amount, iso_currency_code, request.Date, request.Pending, request.MerchantName, request.PaymentChannel)
	if err != nil {
		return err
	}

	return nil
}

// Get all of a user's transactions
func (s *storage) GetTransactions(uid string) ([]api.NewTransactionRequest, error) {

}

// Store a new access token for a user
func (s *storage) CreateAccessToken(request api.NewAccessTokenRequest) error {
	sqlStatement := `
		INSERT INTO plaid_items (user_id, access_token, plaid_item_id)
		VALUES ($1, $2, $3)
		`
	_, err := s.db.Exec(sqlStatement, request.UID, request.AccessToken, request.ItemId)
	if err != nil {
		return err
	}
	return nil
}
