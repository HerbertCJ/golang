package model

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type PostProduct struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
