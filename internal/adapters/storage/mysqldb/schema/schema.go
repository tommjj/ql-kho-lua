package schema

import (
	"database/sql"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"gorm.io/gorm"
)

type User struct {
	ID                   int             `gorm:"primaryKey;autoIncrement"`
	Name                 string          `gorm:"type:VARCHAR(32);not null"`
	Email                string          `gorm:"type:VARCHAR(320);uniqueIndex;not null"`
	Phone                string          `gorm:"type:VARCHAR(16);not null"`
	Role                 domain.Role     `gorm:"type:VARCHAR(10);not null;default:'member'"`
	Password             string          `gorm:"type:VARCHAR(320);not null"`
	Key                  sql.NullString  `gorm:"type:VARCHAR(320)"`
	DeletedAt            gorm.DeletedAt  `gorm:"index"`
	AuthorizedWarehouses []*Warehouse    `gorm:"many2many:authorized"`
	ExportInvoices       []ExportInvoice `gorm:"foreignKey:UserID"`
	ImportInvoices       []ImportInvoice `gorm:"foreignKey:UserID"`
}

type Warehouse struct {
	ID              int             `gorm:"primaryKey;autoIncrement"`
	Name            string          `gorm:"type:VARCHAR(255);uniqueIndex;not null"`
	Location        string          `gorm:"type:VARCHAR(50);not null"`
	Capacity        int             `gorm:"type:INTEGER;not null"`
	Image           string          `gorm:"type:VARCHAR(255);not null"`
	DeletedAt       gorm.DeletedAt  `gorm:"index"`
	AuthorizedUsers []*User         `gorm:"many2many:authorized;"`
	ExportInvoices  []ExportInvoice `gorm:"foreignKey:WarehouseID"`
	ImportInvoices  []ImportInvoice `gorm:"foreignKey:WarehouseID"`
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
	Phone          string          `gorm:"type:VARCHAR(16);not null"`
	Address        string          `gorm:"type:VARCHAR(255);not null"`
	DeletedAt      gorm.DeletedAt  `gorm:"index"`
	ExportInvoices []ExportInvoice `gorm:"foreignKey:CustomerID"`
	ImportInvoices []ImportInvoice `gorm:"foreignKey:CustomerID"`
}

type ExportInvoice struct {
	ID          int                   `gorm:"primaryKey;autoIncrement"`
	WarehouseID int                   `gorm:"not null;index"`
	CustomerID  int                   `gorm:"not null"`
	UserID      int                   `gorm:"not null"`
	TotalPrice  float64               `gorm:"not null"`
	CreatedAt   time.Time             ``
	Warehouse   Warehouse             `gorm:"foreignKey:WarehouseID"`
	Customer    Customer              `gorm:"foreignKey:CustomerID"`
	User        User                  `gorm:"foreignKey:UserID"`
	Details     []ExportInvoiceDetail `gorm:"foreignKey:InvoiceID"`
}

type ExportInvoiceDetail struct {
	InvoiceID int     `gorm:"primaryKey"`
	RiceID    int     `gorm:"primaryKey"`
	Price     float64 `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Rice      Rice    `gorm:"foreignKey:RiceID"`
}

type ImportInvoice struct {
	ID          int                   `gorm:"primaryKey;autoIncrement"`
	WarehouseID int                   `gorm:"not null;index"`
	CustomerID  int                   `gorm:"not null"`
	UserID      int                   `gorm:"not null"`
	TotalPrice  float64               `gorm:"not null"`
	CreatedAt   time.Time             ``
	Warehouse   Warehouse             `gorm:"foreignKey:WarehouseID"`
	Customer    Customer              `gorm:"foreignKey:CustomerID"`
	User        User                  `gorm:"foreignKey:UserID"`
	Details     []ImportInvoiceDetail `gorm:"foreignKey:InvoiceID"`
}

type ImportInvoiceDetail struct {
	InvoiceID int     `gorm:"primaryKey"`
	RiceID    int     `gorm:"primaryKey"`
	Price     float64 `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Rice      Rice    `gorm:"foreignKey:RiceID"`
}
