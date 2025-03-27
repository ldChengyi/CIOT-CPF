package dto

type ProductDTO struct {
	ID            int64  `json:"id" xml:"id" form:"id" query:"id"`
	ProductKey    string `json:"product_key" xml:"product_key" form:"product_key" query:"product_key"`
	ProductName   string `json:"product_name" xml:"product_name" form:"product_name" query:"product_name"`
	ProductSecret string `json:"product_secret" xml:"product_secret" form:"product_secret" query:"product_secret"`
	Desc          string `json:"desc" xml:"desc" form:"desc" query:"desc"`
}
