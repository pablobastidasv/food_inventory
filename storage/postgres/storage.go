package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/pablobastidasv/fridge_inventory/storage"
	"github.com/pablobastidasv/fridge_inventory/types"
)

type PostgresStore struct {
	db *sql.DB
}

func New(db *sql.DB) storage.Store {
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
	query := "insert into products(id, name, category_code) values ($1, $2, $3)"
	if _, err := s.db.ExecContext(c, query, p.Id, p.Name, p.Category.Code); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) ListProducts(c context.Context) ([]types.Product, error) {
	query := "select p.id, p.name, c.code, c.name from products p join categories c on p.category_code = c.code"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prods := []types.Product{}
	for rows.Next() {
		var prod types.Product

		if err := rows.Scan(
			&prod.Id,
			&prod.Name,
			&prod.Category.Code,
			&prod.Category.Name,
		); err != nil {
			return nil, err
		}
		prods = append(prods, prod)
	}
	return prods, nil
}

func (s *PostgresStore) ListCategories(c context.Context) ([]types.Category, error) {
	query := "select c.code, c.name, p.code, p.name from categories c left join categories p on c.parent = p.code"
	rows, err := s.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []types.Category{}
	for rows.Next() {
		var category types.Category
		var parentCode *string
		var parentName *string

		if err := rows.Scan(
			&category.Code,
			&category.Name,
			&parentCode,
			&parentName,
		); err != nil {
			return nil, err
		}

		if parentCode != nil {
			category.Parent = &types.Category{
				Code: *parentCode,
				Name: *parentName,
			}
		}

		categories = append(categories, category)
	}

	return categories, nil
}
