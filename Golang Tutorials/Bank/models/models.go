package models

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// InitializeDB initializes the database connection and performs migrations
func InitializeDB() {
	dsn := "root:Forcepoint@2024@tcp(127.0.0.1:3306)/gobankapp?charset=utf8mb4&parseTime=True&loc=Local" // Update this with your actual DB credentials
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automatically migrate schema for all models
	DB.AutoMigrate(&Bank{})        // Banks table first
	DB.AutoMigrate(&User{})        // Users table second
	DB.AutoMigrate(&Account{})     // Accounts table third, after banks and users
	DB.AutoMigrate(&LedgerEntry{}) // Then ledger entry
	DB.AutoMigrate(&Transaction{}) // Finally transaction
}

// User struct
type User struct {
	gorm.Model
	Username   string     `gorm:"unique;not null" json:"username"`
	Password   string     `gorm:"not null" json:"-"`
	FirstName  string     `gorm:"not null" json:"firstName"`
	LastName   string     `gorm:"not null" json:"lastName"`
	IsAdmin    bool       `gorm:"default:false" json:"isAdmin"`
	IsActive   bool       `gorm:"default:true" json:"isActive"`
	IsCustomer bool       `gorm:"default:false"`
	Accounts   []*Account `gorm:"foreignKey:CustomerID;references:ID" json:"accounts"` //constraint:OnUpdate:CASCADE,OnDelete:CASCADE  -- if a User is deleted, their accounts are also deleted
}

// Account struct - Ensure this is defined before being used in User struct
type Account struct {
	gorm.Model
	CustomerID uint           `gorm:"not null" json:"customer_id"`          // Foreign Key to User table
	BankID     uint           `gorm:"not null" json:"bank_id"`              // Foreign Key to Bank table
	Balance    float64        `gorm:"not null;default:1000" json:"balance"` // Initial balance default is 1000
	IsActive   bool           `gorm:"not null;default:true" json:"is_active"`
	Passbook   []*Transaction `gorm:"foreignKey:AccountID;references:ID" json:"passbook"`
}

// Bank struct
type Bank struct {
	gorm.Model
	FullName     string         `gorm:"not null" json:"full_name"`
	Abbreviation string         `gorm:"size:5;not null" json:"abbreviation"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	Accounts     []*Account     `gorm:"foreignKey:BankID;references:ID"`
	Ledger       []*LedgerEntry `gorm:"foreignKey:BankID;references:ID;foreignKey:CorrespondingBankID;references:ID" json:"ledger"`
}

// LedgerEntry struct
type LedgerEntry struct {
	gorm.Model
	BankID              uint    `gorm:"not null" json:"bank_id"` // Foreign key to Bank
	CorrespondingBankID uint    `gorm:"not null" json:"corresponding_bank_id"`
	Amount              float64 `gorm:"not null" json:"amount"`
	EntryType           string  `gorm:"not null" json:"entry_type"` // Lending or Receiving
}

// Transaction struct
type Transaction struct {
	gorm.Model
	AccountID              uint      `gorm:"not null" json:"account_id"`       // Foreign key to Account
	TransactionType        string    `gorm:"not null" json:"transaction_type"` // "debit" or "credit"
	Amount                 float64   `gorm:"not null" json:"amount"`
	NewBalance             float64   `json:"new_balance"`                     // Balance after the transaction
	Time                   time.Time `gorm:"autoCreateTime" json:"timestamp"` // Timestamp of the transaction
	CorrespondingBankID    int       `json:"correspondingBankId"`
	CorrespondingAccountID uint      `json:"corresponding_account_id"` // Foreign key to another Account
}
