package dto

import "point-of-sale/app/model"

type CreateOrderRequest struct {
	Name        string               `json:"name"`
	OrderOption string               `json:"order_option"`
	Note        string               `json:"note,omitempty"`
	TableNumber int                  `json:"number_table,omitempty"`
	Payment     string               `json:"payment"`
	Items       []CreateItemsRequest `json:"items"`
	User        model.User
}

type CreateItemsRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type SearchCategoryRequest struct {
	Name string `json:"name"`
}

type SearchRequest struct {
	Keyword string `json:"keyword"`
}

type CreateTokenCCRequest struct {
}
