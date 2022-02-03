package tik_lib

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

var usersQueries = []string{
	`create table if not exists Users(
		id text,
		first_name text,
		last_name  text,
		username text,
		phone_number text,
		email text,
		password text,
		avatar_url text,
		email_verified bool,
		phone_number_verified bool,
		type_of_user text,
		primary key(id)
	);`,
}

type userStore struct {
	db *sql.DB
}

func NewPostgreUserStore(cfg PostgresConfig) (UserStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range usersQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &userStore{db: db}
	return store, nil
}

func (u *userStore) Create(user *User) (*User, error) {
	result, err := u.db.Exec(
		"INSERT INTO Users "+
			"(id, first_name, last_name, username, phone_number, email, password, avatar_url, email_verified, phone_number_verified, type_of_user) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		user.Id, user.FirstName, user.LastName, user.Username,
		user.PhoneNumber, user.Email, user.Password, user.AvatarUrl,
		user.EmailVerified, user.PhoneNumberVerified, user.TypeOfUser.ToString(),
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateUserUnknown
	}
	return user, nil
}

func (u *userStore) Update(user *UserUpdate) (*User, error) {
	q := "update users set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if user.FirstName != nil {
		cnt++
		parts = append(parts, "first_name = $"+strconv.Itoa(cnt))
		values = append(values, user.FirstName)
	}
	if user.LastName != nil {
		cnt++
		parts = append(parts, "last_name = $"+strconv.Itoa(cnt))
		values = append(values, user.LastName)
	}
	if user.Username != nil {
		cnt++
		parts = append(parts, "username = $"+strconv.Itoa(cnt))
		values = append(values, user.Username)
	}
	if user.PhoneNumber != nil {
		cnt++
		parts = append(parts, "phone_number = $"+strconv.Itoa(cnt))
		values = append(values, user.PhoneNumber)
	}
	if user.Email != nil {
		cnt++
		parts = append(parts, "email = $"+strconv.Itoa(cnt))
		values = append(values, user.Email)
	}
	if user.Password != nil {
		cnt++
		parts = append(parts, "password = $"+strconv.Itoa(cnt))
		values = append(values, user.Password)
	}
	if user.AvatarUrl != nil {
		cnt++
		parts = append(parts, "avatar_url = $"+strconv.Itoa(cnt))
		values = append(values, user.AvatarUrl)
	}
	if user.EmailVerified != nil {
		cnt++
		parts = append(parts, "email_verified = $"+strconv.Itoa(cnt))
		values = append(values, user.EmailVerified)
	}
	if user.PhoneNumberVerified != nil {
		cnt++
		parts = append(parts, "phone_number_verified = $"+strconv.Itoa(cnt))
		values = append(values, user.PhoneNumberVerified)
	}
	if user.TypeOfUser != nil {
		cnt++
		parts = append(parts, "type_of_user = $"+strconv.Itoa(cnt))
		values = append(values, user.TypeOfUser.ToString())
	}

	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	q = q + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, user.Id)
	result, err := u.db.Exec(q, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrUserNotFound
	}
	return u.Get(user.Id)
}

func (u *userStore) Delete(id string) error {
	result, err := u.db.Exec("delete from users where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrUserNotFound
	}
	return nil
}

func (u *userStore) Get(id string) (*User, error) {
	user := &User{}
	userType := ""
	err := u.db.QueryRow("select "+
		"id, first_name, last_name, username, phone_number, email, password, avatar_url, email_verified, phone_number_verified, type_of_user "+
		"from users where id = $1 limit 1", id).
		Scan(
			&user.Id, &user.FirstName,
			&user.LastName, &user.Username,
			&user.PhoneNumber, &user.Email,
			&user.Password, &user.AvatarUrl,
			&user.EmailVerified, &user.PhoneNumberVerified,
			&userType)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	user.TypeOfUser = ToUserType(userType)
	return user, nil
}

func (u *userStore) List(typeOfUser string) ([]User, error) {
	users := []User{}
	var values []interface{}
	values = append(values, typeOfUser)
	q := "select " +
		"id, first_name, last_name, username, phone_number, email, password, avatar_url, email_verified, phone_number_verified, type_of_user " +
		"from users where type_of_user = $1"
	//cnt := 1
	rows, err := u.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		userType := ""
		err = rows.Scan(
			&user.Id, &user.FirstName,
			&user.LastName, &user.Username,
			&user.PhoneNumber, &user.Email,
			&user.Password, &user.AvatarUrl,
			&user.EmailVerified, &user.PhoneNumberVerified,
			&userType)
		if err != nil {
			return nil, err
		}
		user.TypeOfUser = ToUserType(userType)
		users = append(users, user)
	}
	return users, nil
}
