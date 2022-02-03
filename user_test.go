package tik_lib

import (
	"fmt"
	"github.com/google/uuid"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

var (
	store       UserStore
	currentUser *User
)

func init() {
	filepath, err := os.Getwd()
	if err != nil {
		panic("main, get rootDir error" + err.Error())
		return
	}
	viper.SetConfigName("main")
	viper.AddConfigPath(filepath + "./config/")
	err = viper.ReadInConfig()
	if err != nil {
		panic("main, fatal error while reading config file: " + err.Error())
		return
	}
	cfg := PostgresConfig{
		Host:     viper.GetString("db.primary.host"),
		Port:     5432,
		User:     viper.GetString("db.primary.user"),
		Password: viper.GetString("db.primary.pass"),
		Database: viper.GetString("db.primary.name"),
		Params:   viper.GetString("db.primary.param"),
	}
	st, err := NewPostgreUserStore(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	store = st
	//Db connection
}

func TestUserStore_Create(t *testing.T) {
	filepath, err := os.Getwd()
	if err != nil {
		panic("main, get rootDir error" + err.Error())
		return
	}
	viper.SetConfigName("user_create")
	viper.AddConfigPath(filepath + "./config/")
	err = viper.ReadInConfig()
	if err != nil {
		panic("main, fatal error while reading config file: " + err.Error())
		return
	}
	hashPass, err := setdata_common.HashPassword(viper.GetString("user.password"))
	if err != nil {
		t.Error(err)
		return
	}
	userTest := &User{
		Id:                  uuid.New().String(),
		FirstName:           viper.GetString("user.first_name"),
		LastName:            viper.GetString("user.last_name"),
		Username:            viper.GetString("user.username"),
		PhoneNumber:         viper.GetString("user.phone_number"),
		Email:               viper.GetString("user.email"),
		Password:            hashPass,
		AvatarUrl:           viper.GetString("user.avatar_url"),
		EmailVerified:       viper.GetBool("user.email_verified"),
		PhoneNumberVerified: viper.GetBool("user.phone_number_verified"),
		TypeOfUser:          ToUserType(viper.GetString("user.type_of_user")),
	}
	newUser, err := store.Create(userTest)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("after user create", newUser)
	currentUser = newUser
}

func TestUserStore_Update(t *testing.T) {
	filepath, err := os.Getwd()
	if err != nil {
		panic("main, get rootDir error" + err.Error())
		return
	}
	viper.SetConfigName("user_update")
	viper.AddConfigPath(filepath + "./config/")
	err = viper.ReadInConfig()
	if err != nil {
		panic("main, fatal error while reading config file: " + err.Error())
		return
	}
	firstName := viper.GetString("user.first_name")
	userTypeString := viper.GetString("user.type_of_user")
	if !IsUserTypeExist(userTypeString) {
		t.Error("not correct type of user")
		return
	}
	userType := ToUserType(userTypeString)
	userTest := &UserUpdate{
		Id:         currentUser.Id,
		FirstName:  &firstName,
		TypeOfUser: &userType,
	}
	updated, err := store.Update(userTest)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("after update ", updated)
}

func TestUserStore_List(t *testing.T) {
	filepath, err := os.Getwd()
	if err != nil {
		panic("main, get rootDir error" + err.Error())
		return
	}
	viper.SetConfigName("user_list")
	viper.AddConfigPath(filepath + "./config/")
	err = viper.ReadInConfig()
	if err != nil {
		panic("main, fatal error while reading config file: " + err.Error())
		return
	}
	userTypeString := viper.GetString("user.type_of_user")
	if !IsUserTypeExist(userTypeString) {
		t.Error("not correct type of user")
		return
	}
	users, err := store.List(userTypeString)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("after list ", users)
}

func TestUserStore_Delete(t *testing.T) {
	err := store.Delete(currentUser.Id)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("after delete id:", currentUser.Id)
}
