package schema

import (
	"time"

	"gorm.io/gorm"
)

// ---- Start Entities ----
type CommonColumns struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type User struct {
	CommonColumns
	Email     string `json:"email" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type Session struct {
	CommonColumns
	UserID       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Expense struct {
	CommonColumns
	UserID  uint        `json:"user_id" gorm:"foreign_key:'ID;references:User;constraint:OnUpdate:CASCADE,OnDelete:CASCADE'"`
	Amount  float64     `json:"amount"`
	Note    string      `json:"note"`
	DueDate time.Time   `json:"due_date"`
	Type    ExpenseType `json:"type" sql:"type:ENUM('food', 'transport', 'entertainment', 'health_and_fitness', 'shopping', 'bills_and_utilities', 'travel', 'groceries', 'other')"`
}

// ---- End Entities ----

// ---- Start Enums ----
type ExpenseType string

const (
	Food              ExpenseType = "food"
	Transport         ExpenseType = "transport"
	Entertainment     ExpenseType = "entertainment"
	HealthAndFitness  ExpenseType = "health_and_fitness"
	Shopping          ExpenseType = "shopping"
	BillsAndUtilities ExpenseType = "bills_and_utilities"
	Travel            ExpenseType = "travel"
	Groceries         ExpenseType = "groceries"
	Other             ExpenseType = "other"
)

// ---- End Enums ----

func DBMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Session{}, &Expense{})
	if err != nil {
		return err
	}

	return nil
}
