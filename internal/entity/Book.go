package entity

type Author struct {
	Name    string
	Surname string
}

type Book struct {
	Name         string
	PackageCount int
	Author
}
