package entity

type Product struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type ProductDto struct {
	Id    *int64  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
