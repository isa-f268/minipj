package model

import "time"

type Users struct {
	User_id      uint   `gorm:"primaryKey" json:"user_id"`
	Email        string `gorm:"unique;not null" json:"email"`
	Password     string `gorm:"not null" json:"password"`
	Name         string `json:"name"`
	TotalBalance int    `gorm:"default:0" json:"total_balance"`

	Topups   []Topup    `gorm:"foreignKey:User_id" json:"topups,omitempty"`
	Orders   []Orders   `gorm:"foreignKey:User_id" json:"orders,omitempty"`
	Payments []Payments `gorm:"foreignKey:User_id" json:"payment,omitempty"`
}

type Books struct {
	Book_id       uint   `gorm:"primaryKey" json:"book_id"`
	Name          string `gorm:"not null" json:"name"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Price_per_day int    `gorm:"not null" json:"price_per_day"`

	Orders []Orders `gorm:"foreignKey:Book_id" json:"orders,omitempty"`
}

type Topup struct {
	Topup_id  uint      `gorm:"primaryKey" json:"topup_id"`
	User_id   int       `gorm:"not null" json:"user_id"`
	Balance   int       `gorm:"not null" json:"balance"`
	TopupDate time.Time `json:"topup_date"`
}

type Orders struct {
	Order_id  uint      `gorm:"primaryKey" json:"order_id"`
	User_id   int       `gorm:"not null" json:"user_id"`
	Book_id   int       `gorm:"not null" json:"book_id"`
	TotalDays int       `json:"total_days"`
	Rent_date time.Time `json:"rent_date"`
	Finished  bool      `gorm:"default:false" json:"finished"`

	Payment Payments `gorm:"foreignKey:Order_id" json:"payment,omitempty"`
}

type Payments struct {
	Payment_id uint      `gorm:"primaryKey" json:"payment_id"`
	Order_id   int       `gorm:"unique;not null" json:"order_id"`
	User_id    int       `gorm:"unique;not null" json:"user_id"`
	Amount     int       `gorm:"not null" json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `gorm:"default:false" json:"status"`
}
