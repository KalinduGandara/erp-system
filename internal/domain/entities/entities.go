package entities

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Customer struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID          uint   `gorm:"primarykey"`
	Code        string `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Invoice struct {
	ID          uint   `gorm:"primarykey"`
	Number      string `gorm:"unique;not null"`
	CustomerID  uint   `gorm:"not null"`
	Customer    Customer
	TotalAmount float64
	Items       []InvoiceItem
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type InvoiceItem struct {
	ID         uint `gorm:"primarykey"`
	InvoiceID  uint `gorm:"not null"`
	ProductID  uint `gorm:"not null"`
	Product    Product
	Quantity   int     `gorm:"not null"`
	UnitPrice  float64 `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
}
