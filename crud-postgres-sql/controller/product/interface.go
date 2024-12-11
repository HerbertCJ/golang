package controller

import "net/http"

type IProduct interface {
	GetAll(w http.ResponseWriter, r *http.Request) error
	GetById(w http.ResponseWriter, r *http.Request) error
	Create(w http.ResponseWriter, r *http.Request) error
	Delete(w http.ResponseWriter, r *http.Request) error
	Update(w http.ResponseWriter, r *http.Request) error
}
