package tik_lib

type File struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FileUrl     string `json:"file_url"`
	FileType    FileType `json:"file_type"`
}

type FileUpdate struct {
	Id          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	FileUrl     *string `json:"file_url"`
	FileType    *FileType `json:"file_type"`
}

