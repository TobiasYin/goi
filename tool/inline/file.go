package inline

type File struct {
	IsDir bool
	IsFile bool
	Name string
	Content string
	Children []File
}

var Root File
