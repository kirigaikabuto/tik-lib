package tik_lib

type UserType string

var (
	Shop  UserType = "shop"
	Buyer UserType = "buyer"
)

var (
	userTypeToString = map[UserType]string{
		Shop:  "shop",
		Buyer: "buyer",
	}
	stringToUserTypeType = map[string]UserType{
		"shop":  Shop,
		"buyer": Buyer,
	}
)

func (c UserType) ToString() string {
	return userTypeToString[c]
}

func ToUserType(s string) UserType {
	return stringToUserTypeType[s]
}

func IsUserTypeExist(s string) bool {
	userTypes := []string{"shop", "buyer"}
	for _, v := range userTypes {
		if v == s {
			return true
		}
	}
	return false
}
