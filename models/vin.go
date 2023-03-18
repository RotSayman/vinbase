package models

type Vin struct {
	Name  string `json:"name"`
	VinID string `gorm:"primaryKey" json:"vin_id"`
}
