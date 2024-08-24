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

type InventoryItem struct {
	Id      string
	Product Product
	Amount int
}
