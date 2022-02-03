package tik_lib

type User struct {
	Id                  string `json:"id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Username            string `json:"username"`
	PhoneNumber         string `json:"phone_number"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	AvatarUrl           string `json:"avatar_url"`
	EmailVerified       bool   `json:"email_verified"`
	PhoneNumberVerified bool   `json:"phone_number_verified"`
	TypeOfUser          string `json:"type_of_user"`
}

type UserUpdate struct {
	Id                  string  `json:"id"`
	FirstName           *string `json:"first_name"`
	LastName            *string `json:"last_name"`
	Username            *string `json:"username"`
	PhoneNumber         *string `json:"phone_number"`
	Email               *string `json:"email"`
	Password            *string `json:"password"`
	AvatarUrl           *string `json:"avatar_url"`
	EmailVerified       *bool   `json:"email_verified"`
	PhoneNumberVerified *bool   `json:"phone_number_verified"`
	TypeOfUser          *string `json:"type_of_user"`
}
