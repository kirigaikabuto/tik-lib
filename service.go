package tik_lib

type service struct {
	userStore UserStore
	fileStore FileStore
}

type Service interface {
	//users
	CreateUser(cmd *CreateUserCommand) (*User, error)
	UpdateUser(cmd *UpdateUserCommand) (*User, error)
	GetUserById(cmd *GetUserByIdCommand) (*User, error)
	ListUsers(cmd *ListUserCommand) ([]User, error)
	DeleteUser(cmd *DeleteUserCommand) error
	GetUserByPhoneNumber(cmd *GetUserByPhoneNumberCommand) (*User, error)

	//files
	CreateFile(cmd *CreateFileCommand) (*File, error)
	UpdateFile(cmd *UpdateFileCommand) (*File, error)
	GetFileById(cmd *GetFileByIdCommand) (*File, error)
	ListFiles(cmd *ListFilesCommand) ([]File, error)
	Delete(cmd *DeleteFileCommand) error
}

func NewService(u UserStore, f FileStore) Service {
	return &service{userStore: u, fileStore: f}
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

func (s *service) CreateFile(cmd *CreateFileCommand) (*File, error) {
	return s.fileStore.Create(cmd.File)
}

func (s *service) UpdateFile(cmd *UpdateFileCommand) (*File, error) {
	return s.fileStore.Update(cmd.FileUpdate)
}

func (s *service) GetFileById(cmd *GetFileByIdCommand) (*File, error) {
	return s.fileStore.Get(cmd.Id)
}

func (s *service) ListFiles(cmd *ListFilesCommand) ([]File, error) {
	return s.fileStore.List()
}

func (s *service) Delete(cmd *DeleteFileCommand) error {
	return s.fileStore.Delete(cmd.Id)
}
