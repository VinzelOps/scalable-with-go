package models

import "time"

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Paket        []Paket    `gorm:"foreignKey:OrderID" json:"paket"`
}
