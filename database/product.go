package database

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var ProductList = []Product{
	{
		ID:    1,
		Name:  "Product1",
		Price: 45.66,
	},
	{
		ID:    2,
		Name:  "Product2",
		Price: 45.66,
	},
	{
		ID:    3,
		Name:  "Product3",
		Price: 45.66,
	},
}