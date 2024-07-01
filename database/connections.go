package database

import (
	"context"
	"database/sql"
)

type connections struct {
	db *sql.DB
}

func (c *connections) NewTransaction(ctx context.Context, readOnly bool) (tx Transaction, err error) {

	sqlTx, err := c.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  readOnly,
	})
	if err != nil {
		return nil, err
	}

	t := &transaction{tx: sqlTx}

	return t, nil
}
