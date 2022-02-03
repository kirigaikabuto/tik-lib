package tik_lib

type UserStore interface {
	Create(user *User) (*User, error)
	Update(user *UserUpdate) (*User, error)
	Delete(id string) error
	Get(id string) (*User, error)
	List(typeOfUser string) ([]User, error)
}
