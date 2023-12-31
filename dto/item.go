package dto

import "time"

type NewItemRequest struct {
	ItemCode    string `json:"itemCode" example:"indomiegr"`
	Description string `json:"description" example:"Indomie Goreng"`
	Quantity    int    `json:"quantity" example:"2"`
}

type OrderWithItems struct {
	OrderId      int               `json:"orderId"`
	CustomerName string            `json:"customerName"`
	OrderedAt    time.Time         `json:"orderedAt"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	Items        []GetItemResponse `json:"items"`
}
