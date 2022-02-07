package tik_lib

type FileType string

var (
	Video FileType = "video"
	Image FileType = "image"
)

var (
	fileTypeToString = map[FileType]string{
		Video: "video",
		Image: "image",
	}
	stringToFileType = map[string]FileType{
		"video": Video,
		"image": Image,
	}
)

func (c FileType) ToString() string {
	return fileTypeToString[c]
}

func ToFileType(s string) FileType {
	return stringToFileType[s]
}

func IsFileTypeExist(s string) bool {
	fileTypes := []string{"video", "image"}
	for _, v := range fileTypes {
		if v == s {
			return true
		}
	}
	return false
}
