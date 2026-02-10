package repo

import (
	"fmt"
	"go-server/domain"
	"go-server/user"

	"github.com/jmoiron/sqlx"
)


type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func (r *userRepo) Create(user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		) VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_shop_owner
		)
		RETURNING id, created_at, updated_at
	`

	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println("Error happend", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID= userID

	return user, nil
}
func (r *userRepo) Find(email string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			is_shop_owner,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}


// func (r *userRepo) List() ([]*User, error) {
// 	return r.users, nil
// }
// func (r *userRepo) Delete(productId int) error {
// 	var temList []*User

// 	for _, product := range r.users {
// 		if product.ID != productId {
// 			temList = append(temList, product)
// 		}
// 	}

// 	r.users = temList

// 	return nil
// }
// func (r *userRepo) Update(user User) (*User, error) {
// 	for idx, p := range r.users {
// 		if p.ID == user.ID {
// 			r.users[idx] = &user
// 		}
// 	}
// 	return &user, nil
// }

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

