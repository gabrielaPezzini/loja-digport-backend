package model

type Cart struct {
	UserId       string
	TotalPrice   float64
	Id           string
	InfosProduct []InfosProduct
}

type InfosProduct struct {
	ProductId       string
	ProductQuantity int
}
