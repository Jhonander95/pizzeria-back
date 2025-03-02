package models

import (
	"time"
)

type Order struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Products  []OrderProduct `json:"products" gorm:"foreignKey:OrderID"`
	Total     float64        `json:"total"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
}

type OrderProduct struct {
	OrderID   uint    `json:"order_id" gorm:"primaryKey"`
	ProductID uint    `json:"product_id" gorm:"primaryKey"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
}

/* type Product struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
} */
