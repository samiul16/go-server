package repo

import (
	"database/sql"
	"fmt"
	"go-server/domain"
	"go-server/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func (r *productRepo) Create(p *domain.Product) (*domain.Product, error) {
  query := `
   INSERT INTO products(
    title,
	description,
	price,
	img_url  
   ) VALUES (
	$1,
	$2,
	$3,
	$4
	)

	RETURNING id
  `

  row := r.db.QueryRow(query, p.Title, p.Description,p.Price, p.ImgUrl)
  err:= row.Scan(&p.ID)

  if err !=nil {
	fmt.Println("Some thing wrong on scanning row after insert")
  }

  return p, nil

}
func (r *productRepo) Get(id int) (*domain.Product, error) {
	query := `
	SELECT * from products
	WHERE id = $1`

	product := domain.Product{}

	err := r.db.Get(&product, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil,nil
		}
	}

	return &product,nil
}

func (r *productRepo) List() ([]*domain.Product, error) {
	query := `SELECT * from products`

	var prdList []*domain.Product

	err := r.db.Select(&prdList, query)

	if err != nil {
			return nil,nil
	}

	return prdList,nil
}
func (r *productRepo) Delete(productId int) error {
	query := `DELETE FROM products WHERE id = ?`

	_, err := r.db.Exec(query, productId)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) Update(product *domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET title = ?, price = ?, description = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		product.Title,
		product.Price,
		product.Description,
		product.ID,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepo) Find(productId int) (*domain.Product, error) {
	query := `
	SELECT * from products
	WHERE id = $1`

	product := domain.Product{}

	err := r.db.Get(&product, query, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil,nil
		}
	}

	return &product,nil
}


func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}