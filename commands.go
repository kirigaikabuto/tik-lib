package tik_lib

type CreateUserCommand struct {
	*User
}

func (cmd *CreateUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).CreateUser(cmd)
}

type UpdateUserCommand struct {
	*UserUpdate
}

func (cmd *UpdateUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).UpdateUser(cmd)
}

type DeleteUserCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteUserCommand) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(Service).DeleteUser(cmd)
}

type ListUserCommand struct {
	TypeOfUser string `json:"type_of_user"`
}

func (cmd *ListUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).ListUsers(cmd)
}

type GetUserByIdCommand struct {
	Id string `json:"id"`
}

func (cmd *GetUserByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetUserById(cmd)
}

type GetUserByPhoneNumberCommand struct {
	PhoneNumber string `json:"phone_number"`
}

func (cmd *GetUserByPhoneNumberCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).GetUserByPhoneNumber(cmd)
}
