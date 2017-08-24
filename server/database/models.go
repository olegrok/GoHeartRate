package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User stores information about user: ID, username, password hash, password salt and creation time
type User struct {
	ID        uint   `gorm:"primary_key; AUTO_INCREMENT"`
	Username  string `gorm:"not null;unique;size:64"`
	Password  string `gorm:"not null"`
	Salt      string `gorm:"not null"`
	CreatedAt time.Time
}

// BeforeCreate initializes "create_at" field in User object before creation new user in database
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedAt", time.Now())
}

// UserSession stores information about user's session: ID, UserID, Auth token, creation and expiration time
type UserSession struct {
	ID        uint64 `gorm:"primary_key; AUTO_INCREMENT"`
	User      User   `gorm:"ForeignKey:UserID"`
	UserID    uint   `gorm:"not null"`
	Token     string `gorm:"not null"`
	CreatedAt time.Time
	ExpireAt  time.Time `gorm:"DEFAULT: NULL"`
}

// BeforeCreate initializes "create_at" field in UserSession object before creation new user's session in database
func (user *UserSession) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedAt", time.Now())
}

// UserResult stores information about user's result: ID, UserID, result and creation time
type UserResult struct {
	ID        uint64 `gorm:"primary_key; AUTO_INCREMENT"`
	User      User   `gorm:"ForeignKey:UserID"`
	UserID    uint64 `gorm:"not null"`
	Result    string `gorm:"not null"`
	CreatedAt time.Time
}

// BeforeCreate initializes "create_at" field in UserResult object before creation new user's result in database
func (user *UserResult) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedAt", time.Now())
}
