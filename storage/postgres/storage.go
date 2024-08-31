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
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	if parent != nil {
		parent, err := p.FindCategory(c, *parent)
		if err != nil {
			return nil, err
		}
		category.Parent = parent
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

func (s *PostgresStore) DeleteProduct(ctx context.Context, id string) error {
	query := "delete from products p where p.id = $1"
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) ListInventoryItems(c context.Context) ([]types.InventoryItem, error) {
	query := `
        select 
            ii.id, ii.amount ,
            p.id , p.name,
            c.code, c.name,
            cp.code, cp."name" 
        from inventory_items ii 
        join products p on ii.product_id = p.id 
        join categories c on p.category_code = c.code 
        left join categories cp on c.parent = cp.code
        order by p.name
    `
	rows, err := s.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []types.InventoryItem{}
	for rows.Next() {
		var i types.InventoryItem
		var categoryParentCode *string
		var categoryParentName *string

		if err := rows.Scan(
			&i.Id,
			&i.Amount,
			&i.Product.Id,
			&i.Product.Name,
			&i.Product.Category.Code,
			&i.Product.Category.Name,
			&categoryParentCode,
			&categoryParentName,
		); err != nil {
			return nil, err
		}

		if categoryParentCode != nil {
			i.Product.Category.Parent = &types.Category{
				Code: *categoryParentCode,
				Name: *categoryParentName,
			}
		}

		items = append(items, i)
	}

	return items, nil
}

func (s *PostgresStore) FindInventoryItemById(ctx context.Context, id string) (*types.InventoryItem, error) {
	query := `
        select 
            ii.id, ii.amount ,
            p.id , p.name,
            c.code, c.name,
            cp.code, cp.name 
        from inventory_items ii 
        join products p on ii.product_id = p.id 
        join categories c on p.category_code = c.code 
        left join categories cp on c.parent = cp.code
        where ii.id = $1
    `
	row := s.db.QueryRowContext(ctx, query, id)

	var i types.InventoryItem
	var categoryParentCode *string
	var categoryParentName *string

	if err := row.Scan(
		&i.Id,
		&i.Amount,
		&i.Product.Id,
		&i.Product.Name,
		&i.Product.Category.Code,
		&i.Product.Category.Name,
		&categoryParentCode,
		&categoryParentName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if categoryParentCode != nil {
		i.Product.Category.Parent = &types.Category{
			Code: *categoryParentCode,
			Name: *categoryParentName,
		}
	}

	return &i, nil
}

func (s *PostgresStore) UpdateInventoryItem(c context.Context, id string, amount int) error {
	query := `
    update inventory_items ii
        set amount = $1
    where ii.id = $2
    `
	if _, err := s.db.ExecContext(c, query, amount, id); err != nil {
		return err
	}

	return nil
}
