package controller

import (
	"crud-postgres-sql/model"
	repository "crud-postgres-sql/repository/product"
	"database/sql"
	"encoding/json"
	"net/http"
)

type ProductController struct {
	Db *sql.DB
}

func NewProductController(db *sql.DB) IProduct {
	return &ProductController{Db: db}
}

func (p *ProductController) GetAll(w http.ResponseWriter, r *http.Request) error {
	repository := repository.NewProductRepository(p.Db)

	products, err := repository.GetAll()

	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"data": products,
	}
	json.NewEncoder(w).Encode(response)

	return nil
}

func (p *ProductController) GetById(w http.ResponseWriter, r *http.Request) error {
	repository := repository.NewProductRepository(p.Db)

	id := r.PathValue("id")
	product, err := repository.GetById(id)

	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return err
	}

	if product != (model.Product{}) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	} else {
		http.Error(w, `{"error": "Not Found"}`, http.StatusNotFound)
	}

	return nil
}

func (p *ProductController) Create(w http.ResponseWriter, r *http.Request) error {
	var product model.PostProduct
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, `{"error": "Bad Request"}`, http.StatusBadRequest)
		return err
	}

	repository := repository.NewProductRepository(p.Db)
	err := repository.Create(product)

	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("product created")

	return nil
}

func (p *ProductController) Delete(w http.ResponseWriter, r *http.Request) error {
	repository := repository.NewProductRepository(p.Db)

	id := r.PathValue("id")
	err := repository.Delete(id)

	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("product deleted")

	return nil
}

func (p *ProductController) Update(w http.ResponseWriter, r *http.Request) error {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, `{"error": "Bad Request"}`, http.StatusBadRequest)
		return err
	}

	id := r.PathValue("id")
	repository := repository.NewProductRepository(p.Db)
	product, err := repository.Update(product, id)

	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return err
	}

	if product != (model.Product{}) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("product created")
	} else {
		http.Error(w, `{"error": "Not Found"}`, http.StatusNotFound)
	}

	return nil
}
