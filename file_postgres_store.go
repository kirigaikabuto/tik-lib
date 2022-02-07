package tik_lib

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
)

var filesQueries = []string{
	`create table if not exists Files(
		id text,
		name text,
		description text,
		file_url text,
		file_type text,
		primary key(id)
	);`,
}

type fileStore struct {
	db *sql.DB
}

func NewPostgresFileStore(cfg PostgresConfig) (FileStore, error) {
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
	store := &fileStore{db: db}
	return store, nil
}

func (f *fileStore) Create(file *File) (*File, error) {
	if file.Id == "" {
		file.Id = uuid.New().String()
	}
	result, err := f.db.Exec(
		"INSERT INTO Files "+
			"(id, name, description, file_url, file_type) "+
			"VALUES ($1, $2, $3, $4, $5)",
		file.Id, file.Name, file.Description, file.FileUrl, file.FileType,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateFileUnknown
	}
	return file, nil
}

func (f *fileStore) Update(file *FileUpdate) (*File, error) {
	q := "update Files set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if file.Name != nil {
		cnt++
		parts = append(parts, "name = $"+strconv.Itoa(cnt))
		values = append(values, file.Name)
	}
	if file.FileType != nil {
		cnt++
		parts = append(parts, "file_type = $"+strconv.Itoa(cnt))
		values = append(values, file.FileType.ToString())
	}
	if file.Description != nil {
		cnt++
		parts = append(parts, "description = $"+strconv.Itoa(cnt))
		values = append(values, file.Description)
	}
	if file.FileUrl != nil {
		cnt++
		parts = append(parts, "file_url = $"+strconv.Itoa(cnt))
		values = append(values, file.FileUrl)
	}
	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	q = q + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, file.Id)
	result, err := f.db.Exec(q, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrFileNotFound
	}
	return f.Get(file.Id)
}

func (f *fileStore) Get(id string) (*File, error) {
	file := &File{}
	fileType := ""
	err := f.db.QueryRow("select id, name, description, file_url, file_type from Files where id = $1 limit 1", id).
		Scan(&file.Id, &file.Name, &file.Description, &file.FileUrl, &fileType)
	if err == sql.ErrNoRows {
		return nil, ErrFileNotFound
	} else if err != nil {
		return nil, err
	}
	file.FileType = ToFileType(fileType)
	return file, nil
}

func (f *fileStore) Delete(id string) error {
	result, err := f.db.Exec("delete from Files where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrFileNotFound
	}
	return nil
}

func (f *fileStore) List() ([]File, error) {
	files := []File{}
	var values []interface{}
	q := "select id, name, description, file_url, file_type from Files"
	//cnt := 1
	rows, err := f.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		file := File{}
		fileType := ""
		err = rows.Scan(&file.Id, &file.Name, &file.Description, &file.FileUrl, &fileType)
		if err != nil {
			return nil, err
		}
		file.FileType = ToFileType(fileType)
		files = append(files, file)
	}
	return files, nil
}
