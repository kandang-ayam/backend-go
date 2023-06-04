package dto

type CreateProductRequest struct {
	ProductID   string `form:"products_id" `
	Name        string `form:"products_name"`
	CategoryID  int    `form:"products_category"`
	Quantity    int    `form:"products_quantity"`
	Price       int    `form:"products_price"`
	Unit        string `form:"products_unit"`
	Description string `form:"products_description"`
}
