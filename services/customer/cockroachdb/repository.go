package cockroachdb

import (
	"database/sql"
	"github.com/cockroachdb/cockroach-go/crdb"
	"customerapp/services/customer"
	"context"
	"github.com/go-kit/kit/log"

	"errors"
)


var (
	ErrRepository = errors.New("unable to handle request")
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

// New returns a concrete repository backed by CockroachDB
func New(db *sql.DB, logger log.Logger) (customer.Repository, error) {
	// return  repository
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "cockroachdb"),
	}, nil
}



func (repo *repository) CreateCustomer(ctx context.Context, customer customer.Customer) error {

	// Run a transaction to sync the query model.
	err := crdb.ExecuteTx(ctx, repo.db, nil, func(tx *sql.Tx) error {
		return createCustomer(tx, customer)
	})
	if err != nil {
		return err
	}
	return nil
}





func createCustomer(tx *sql.Tx, customer customer.Customer) error {

	// Insert order into the "orders" table.
	sql := `
			INSERT INTO customer (id, customerid,  address1)
			VALUES ($1,$2,$3)`
	_, err := tx.Exec(sql, customer.ID, customer.CustomerID,  customer.Address)
	if err != nil {
		return err
	}
	return nil
}



// Close implements DB.Close
func (repo *repository) Close() error {
	return repo.db.Close()
}

