package dto

import "point-of-sale/app/model"

type CreateOrderRequest struct {
	Name        string               `json:"name"`
	OrderOption string               `json:"order_option"`
	TableNumber int                  `json:"number_table"`
	Payment     string               `json:"payment"`
	Items       []CreateItemsRequest `json:"items"`
	User        model.User
}

type CreateItemsRequest struct {
	ProductID int    `json:"product_id"`
	Note      string `json:"note,omitempty"`
	Quantity  int    `json:"quantity"`
}

type SearchCategoryRequest struct {
	Name string `json:"name"`
}

type SearchRequest struct {
	Keyword string `json:"keyword"`
}

type CreateTokenCCRequest struct {
}
