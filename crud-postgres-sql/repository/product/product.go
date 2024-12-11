package repository

import (
	"crud-postgres-sql/model"
	"database/sql"
)

type ProductRepostory struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) IProduct {
	return &ProductRepostory{Db: db}
}

func (p *ProductRepostory) GetAll() ([]model.Product, error) {
	var products []model.Product
	rows, err := p.Db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var product model.Product
			if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}

	return products, nil
}

func (p *ProductRepostory) GetById(id string) (model.Product, error) {
	var product model.Product
	rows, err := p.Db.Query("SELECT * FROM products WHERE id = $1", id)

	if err != nil {
		return model.Product{}, err
	}

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
				return model.Product{}, err
			}
		}
	}

	return product, nil
}

func (p *ProductRepostory) Create(product model.PostProduct) error {
	stmt, err := p.Db.Prepare("INSERT INTO products (name, price) VALUES ($1, $2)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepostory) Delete(id string) error {
	_, err := p.Db.Exec("DELETE FROM products WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepostory) Update(product model.Product, id string) (model.Product, error) {
	_, err := p.Db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", product.Name, product.Price, id)

	if err != nil {
		return model.Product{}, err
	}

	return p.GetById(id)
}
