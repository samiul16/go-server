package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner string `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(eamil, pass string) (*User, error)
	List() ([]*User, error)
	Delete(userId int) error
	Update(user User) (*User, error)
}

type userRepo struct {
	users []*User
}

func (r *userRepo) Create(user User) (*User, error) {
	user.ID = len(r.users) + 1
	r.users = append(r.users, &user)
	return &user, nil
}
func (r *userRepo) Find(eamil, pass string) (*User, error) {
	for _, user := range r.users {
		if user.Email == eamil && user.Password == pass {
			return user, nil
		}
	}
	return nil, nil
}

func (r *userRepo) List() ([]*User, error) {
	return r.users, nil
}
func (r *userRepo) Delete(productId int) error {
	var temList []*User

	for _, product := range r.users {
		if product.ID != productId {
			temList = append(temList, product)
		}
	}

	r.users = temList

	return nil
}
func (r *userRepo) Update(user User) (*User, error) {
	for idx, p := range r.users {
		if p.ID == user.ID {
			r.users[idx] = &user
		}
	}
	return &user, nil
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}
