package models

type Order struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Products []Product `json:"products" gorm:"many2many:pedido_productos;"`
	Total    float64   `json:"total"`
}
