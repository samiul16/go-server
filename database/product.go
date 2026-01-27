package database

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var productList = []Product{
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

func Get(pId int) *Product {
	for _, product := range productList {
		if product.ID == pId {
			return &product
		}
	}

	return nil
}

func Store(p Product) {
	p.ID = len(productList) + 1
	productList = append(productList, p)
}

func List() []Product {
	return productList
}
