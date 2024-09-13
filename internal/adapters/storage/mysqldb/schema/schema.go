package schema

import (
	"time"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"gorm.io/gorm"
)

type User struct {
	ID                    int             `gorm:"primaryKey;autoIncrement"`
	Name                  string          `gorm:"type:VARCHAR(32);not null"`
	Email                 string          `gorm:"type:VARCHAR(320);uniqueIndex;not null"`
	Phone                 string          `gorm:"type:VARCHAR(11);not null"`
	Role                  domain.Role     `gorm:"type:VARCHAR(10);not null;default:'staff'"`
	Password              string          `gorm:"type:VARCHAR(320);not null"`
	DeletedAt             gorm.DeletedAt  `gorm:"index"`
	AuthorizedStorehouses []*Storehouse   `gorm:"many2many:authorized"`
	ExportInvoices        []ExportInvoice `gorm:"foreignKey:UserID"`
	ImportInvoices        []ImportInvoice `gorm:"foreignKey:UserID"`
}

type Storehouse struct {
	ID              int             `gorm:"primaryKey;autoIncrement"`
	Name            string          `gorm:"type:VARCHAR(255);not null"`
	Location        string          `gorm:"type:VARCHAR(50);not null"`
	Capacity        int             `gorm:"type:INTEGER;not null"`
	DeletedAt       gorm.DeletedAt  `gorm:"index"`
	AuthorizedUsers []*User         `gorm:"many2many:authorized;"`
	ExportInvoices  []ExportInvoice `gorm:"foreignKey:StorehouseID"`
	ImportInvoices  []ImportInvoice `gorm:"foreignKey:StorehouseID"`
}

type Rice struct {
	ID                   int                   `gorm:"primaryKey;autoIncrement"`
	Name                 string                `gorm:"type:VARCHAR(50);not null;uniqueIndex"`
	DeletedAt            gorm.DeletedAt        `gorm:"index"`
	ExportInvoiceDetails []ExportInvoiceDetail `gorm:"foreignKey:RiceID"`
	ImportInvoiceDetails []ImportInvoiceDetail `gorm:"foreignKey:RiceID"`
}

type Customer struct {
	ID             int             `gorm:"primaryKey;autoIncrement"`
	Name           string          `gorm:"type:VARCHAR(255);not null"`
	Email          string          `gorm:"type:VARCHAR(320);not null"`
	Phone          string          `gorm:"type:VARCHAR(11);not null"`
	Address        string          `gorm:"type:VARCHAR(255);not null"`
	ExportInvoices []ExportInvoice `gorm:"foreignKey:CustomerID"`
	ImportInvoices []ImportInvoice `gorm:"foreignKey:CustomerID"`
}

type ExportInvoice struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	StorehouseID int `gorm:"not null"`
	CustomerID   int `gorm:"not null"`
	UserID       int `gorm:"not null"`
	CreatedAt    time.Time
	Storehouse   Storehouse            `gorm:"foreignKey:StorehouseID"`
	Customer     Customer              `gorm:"foreignKey:CustomerID"`
	User         User                  `gorm:"foreignKey:UserID"`
	Details      []ExportInvoiceDetail `gorm:"foreignKey:InvoiceID"`
}

type ExportInvoiceDetail struct {
	InvoiceID     int           `gorm:"primaryKey"`
	RiceID        int           `gorm:"primaryKey"`
	Price         float64       `gorm:"not null"`
	Quantity      int           `gorm:"not null"`
	Rice          Rice          `gorm:"foreignKey:RiceID"`
	ExportInvoice ExportInvoice `gorm:"foreignKey:InvoiceID"`
}

type ImportInvoice struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	StorehouseID int `gorm:"not null"`
	CustomerID   int `gorm:"not null"`
	UserID       int `gorm:"not null"`
	CreatedAt    time.Time
	Storehouse   Storehouse            `gorm:"foreignKey:StorehouseID"`
	Customer     Customer              `gorm:"foreignKey:CustomerID"`
	User         User                  `gorm:"foreignKey:UserID"`
	Details      []ImportInvoiceDetail `gorm:"foreignKey:InvoiceID"`
}

type ImportInvoiceDetail struct {
	InvoiceID     int           `gorm:"primaryKey"`
	RiceID        int           `gorm:"primaryKey"`
	Price         float64       `gorm:"not null"`
	Quantity      int           `gorm:"not null"`
	Rice          Rice          `gorm:"foreignKey:RiceID"`
	ImportInvoice ImportInvoice `gorm:"foreignKey:InvoiceID"`
}
