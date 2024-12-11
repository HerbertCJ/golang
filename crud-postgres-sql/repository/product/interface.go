package repository

import "crud-postgres-sql/model"

type IProduct interface {
	GetAll() ([]model.Product, error)
	GetById(id string) (model.Product, error)
	Create(product model.PostProduct) error
	Delete(id string) error
	Update(product model.Product, id string) (model.Product, error)
}
