package schema

import (
	"time"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

//uuid.UUID `gorm:"type:BINARY(16);primaryKey;default:UNHEX(REPLACE(UUID(), '-', ''))"`

type User struct {
	ID                    int             `gorm:"primaryKey;autoIncrement"`
	Name                  string          `gorm:"type:VARCHAR(32)"`
	Email                 string          `gorm:"type:VARCHAR(320);uniqueIndex"`
	Phone                 string          `gorm:"type:VARCHAR(11)"`
	Rule                  domain.Role     `gorm:"type:VARCHAR(10)"`
	Password              string          `gorm:"type:VARCHAR(320)"`
	AuthorizedStorehouses []*Storehouse   `gorm:"many2many:authorized;"`
	ExportInvoices        []ExportInvoice `gorm:"foreignKey:UserID"`
	ImportInvoices        []ImportInvoice `gorm:"foreignKey:UserID"`
}

type Storehouse struct {
	ID              int             `gorm:"primaryKey;autoIncrement"`
	Name            string          `gorm:"type:VARCHAR(255)"`
	Location        string          `gorm:"type:VARCHAR(50)"`
	Capacity        int             `gorm:"type:INTEGER"`
	AuthorizedUsers []*User         `gorm:"many2many:authorized;"`
	ExportInvoices  []ExportInvoice `gorm:"foreignKey:StorehouseID"`
	ImportInvoices  []ImportInvoice `gorm:"foreignKey:StorehouseID"`
}

type Rice struct {
	ID                   int                   `gorm:"primaryKey;autoIncrement"`
	Name                 string                `gorm:"type:VARCHAR(50)"`
	ExportInvoiceDetails []ExportInvoiceDetail `gorm:"foreignKey:RiceID"`
	ImportInvoiceDetails []ImportInvoiceDetail `gorm:"foreignKey:RiceID"`
}

type Customer struct {
	ID             int             `gorm:"primaryKey;autoIncrement"`
	Name           string          `gorm:"type:VARCHAR(255)"`
	Email          string          `gorm:"type:VARCHAR(320)"`
	Phone          string          `gorm:"type:VARCHAR(11)"`
	Address        string          `gorm:"type:VARCHAR(255)"`
	ExportInvoices []ExportInvoice `gorm:"foreignKey:CustomerID"`
	ImportInvoices []ImportInvoice `gorm:"foreignKey:CustomerID"`
}

type ExportInvoice struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	StorehouseID int
	CustomerID   int
	UserID       int
	CreatedAt    time.Time
	Storehouse   Storehouse            `gorm:"foreignKey:StorehouseID"`
	Customer     Customer              `gorm:"foreignKey:CustomerID"`
	User         User                  `gorm:"foreignKey:UserID"`
	Details      []ExportInvoiceDetail `gorm:"foreignKey:InvoiceID"`
}

type ExportInvoiceDetail struct {
	InvoiceID     int `gorm:"primaryKey"`
	RiceID        int `gorm:"primaryKey"`
	Price         float64
	Quantity      int
	Rice          Rice          `gorm:"foreignKey:RiceID"`
	ExportInvoice ExportInvoice `gorm:"foreignKey:InvoiceID"`
}

type ImportInvoice struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	StorehouseID int
	CustomerID   int
	UserID       int
	CreatedAt    time.Time
	Storehouse   Storehouse            `gorm:"foreignKey:StorehouseID"`
	Customer     Customer              `gorm:"foreignKey:CustomerID"`
	User         User                  `gorm:"foreignKey:UserID"`
	Details      []ImportInvoiceDetail `gorm:"foreignKey:InvoiceID"`
}

type ImportInvoiceDetail struct {
	InvoiceID     int `gorm:"primaryKey"`
	RiceID        int `gorm:"primaryKey"`
	Price         float64
	Quantity      int
	Rice          Rice          `gorm:"foreignKey:RiceID"`
	ImportInvoice ImportInvoice `gorm:"foreignKey:InvoiceID"`
}
