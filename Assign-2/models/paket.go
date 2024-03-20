package models

type Paket struct {
	PaketID      uint   `gorm:"primaryKey" json:"paket_id"`
	PaketCode    string `json:"paket_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
