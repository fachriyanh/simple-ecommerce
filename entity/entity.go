package entity

var ListProduct = make(map[string]*Product)

type Product struct {
	Name  string  `json:"name"`
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var ListOrder []*Order

type Order struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Items int    `json:"items"`
}
