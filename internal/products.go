package models

import (
	"context"
	"time"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) Insert(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int

	stmt := `insert into products (name, description, price, created_at, updated_at)
	values ($1, $2, $3, $4, $5) returning id`

	err := db.QueryRowContext(ctx, stmt,
		product.Name,
		product.Description,
		product.Price,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (p *Product) GetOneById(id int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, price, created_at, updated_at from products where id = $1`

	var product Product
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) Update(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel() // resource leaks

	stmt := `update products set
		name = $1,
		description = $2,
		price = $3,
		updated_at = $5
		where id = $6`

	_, err := db.ExecContext(ctx, stmt,
		product.Name,
		product.Description,
		product.Price,
		time.Now(),
		product.Id,
	)

	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (p *Product) GetAll() ([]*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, price, created_at, updated_at from products order by name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *Product) DeleteByID(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from products where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
