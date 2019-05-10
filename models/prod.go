package models

// ProductsMainData список товаров с основными данными
type ProductsMainData struct {
	Code    int
	InStock int
	Name    string
}

// ProductsOtherData список товаров с остальными данными
type ProductsOtherData struct {
	Price1 float64
	Price2 float64
	Code   int
}

// ProductsFullData список товаров со всеми данными
type ProductsFullData struct {
	Code    int
	InStock int
	Name    string
	Price1  float64
	Price2  float64
}
