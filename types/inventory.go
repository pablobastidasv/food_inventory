package types

type Product struct {
	Id       string
	Name     string
	Category Category
}

type Category struct {
	Code   string
	Name   string
	Parent *Category
}
