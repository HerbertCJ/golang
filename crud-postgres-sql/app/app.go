package app

import (
	controller "crud-postgres-sql/controller/product"
	"crud-postgres-sql/db"
	"database/sql"
	"fmt"
	"net/http"
)

type App struct {
	Db *sql.DB
}

func (a *App) Initialize() {
	mux := http.NewServeMux()
	db := db.ConnectDb()

	productController := controller.NewProductController(db)

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.GetAll(w, r)
		case http.MethodPost:
			productController.Create(w, r)
		default:
			http.Error(w, `{"error": "Method Not Allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.GetById(w, r)
		case http.MethodDelete:
			productController.Delete(w, r)
		case http.MethodPatch:
			productController.Update(w, r)
		default:
			http.Error(w, `{"error": "Method Not Allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
