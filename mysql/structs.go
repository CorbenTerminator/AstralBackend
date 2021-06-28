package mysql

import (
	"time"
)

type product struct {
	ID         uint64  `json:"product_id,omitempty"`
	Name       string  `json:"product_name"`
	Price      float32 `json:"price"`
	CategoryID int     `json:"category_id"`
}

type category struct {
	ID   int    `json:"category_id,omitempty"`
	Name string `json:"category"`
}

type order struct {
	ID         uint64    `json:"order_id,omitempty"`
	UserID uint64    `json:"customer_id"`
	CreatedAt  time.Time `json:"createdAt"`
	StatusID   int       `json:"status_id"`
}

type status struct {
	ID         int    `json:"status_id,omitempty"`
	StatusName string `json:"status"`
}

type orderProducts struct {
	OrderID   uint64 `json:"order_id"`
	ProductID uint64 `json:"product_id"`
}

type User struct {
	UserID   uint64 `json:"user_id,omitempty"`
	Login    string `json:"login"`
	Password string `json:"password"`
}