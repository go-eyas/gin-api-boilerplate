package model

type Product struct {
	ID    uint
	Title string
}

func (Product) TableName() string {
	return "products1234"
}
