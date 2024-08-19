package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/pablobastidasv/fridge_inventory/types"
)

type PostgresStore struct {
	db *sql.DB
}

func New(db *sql.DB) *PostgresStore {
	return &PostgresStore{
		db: db,
	}
}

func (p *PostgresStore) FindCategory(c context.Context, categoryCode string) (*types.Category, error) {
    slog.Debug("FindCategory by code", "code", categoryCode)

	query := "select * from categories where code = $1"
	row := p.db.QueryRowContext(c, query, categoryCode)

	var category types.Category
	var parent *string
	if err := row.Scan(
		&category.Code,
		&category.Name,
		&parent,
	); err != nil {
		// TODO: Control when error is ErrNoRows
		return nil, err
	}

	if parent != nil {
		var err error
		category.Parent, err = p.FindCategory(c, *parent)
		if err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func (s *PostgresStore) SaveProduct(c context.Context, p types.Product) error {
	query := "insert into products(id, name, category_code) valuer ($1, $2, $3)"
	if _, err := s.db.ExecContext(c, query, p.Id, p.Name, p.Category.Code); err != nil {
		return err
	}
	return nil
}
