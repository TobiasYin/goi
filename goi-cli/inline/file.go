package inline

type File struct {
	IsDir bool
	IsFile bool
	Name string
	Content []byte
	Children []File
}

var Root File
