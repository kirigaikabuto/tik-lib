package tik_lib

type service struct {
	userStore UserStore
}

type Service interface {
	//users
	CreateUser(cmd *CreateUserCommand) (*User, error)
	UpdateUser(cmd *UpdateUserCommand) (*User, error)
	GetUserById(cmd *GetUserByIdCommand) (*User, error)
	ListUsers(cmd *ListUserCommand) ([]User, error)
	DeleteUser(cmd *DeleteUserCommand) error
	GetUserByPhoneNumber(cmd *GetUserByPhoneNumberCommand) (*User, error)
}

func NewService(u UserStore) Service {
	return &service{userStore: u}
}

func (s *service) CreateUser(cmd *CreateUserCommand) (*User, error) {
	return s.userStore.Create(cmd.User)
}

func (s *service) UpdateUser(cmd *UpdateUserCommand) (*User, error) {
	return s.userStore.Update(cmd.UserUpdate)
}

func (s *service) GetUserById(cmd *GetUserByIdCommand) (*User, error) {
	return s.userStore.Get(cmd.Id)
}

func (s *service) ListUsers(cmd *ListUserCommand) ([]User, error) {
	return s.userStore.List(cmd.TypeOfUser)
}

func (s *service) DeleteUser(cmd *DeleteUserCommand) error {
	return s.userStore.Delete(cmd.Id)
}

func (s *service) GetUserByPhoneNumber(cmd *GetUserByPhoneNumberCommand) (*User, error) {
	return s.userStore.GetByPhoneNumber(cmd.PhoneNumber)
}
