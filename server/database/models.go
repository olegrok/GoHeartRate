package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key; AUTO_INCREMENT"`
	Username  string `gorm:"not null;unique;size:64"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

type UserSession struct {
	ID        uint64 `gorm:"primary_key; AUTO_INCREMENT"`
	User      User   `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID    uint   `gorm:"not null"`
	Token     string `gorm:"not null"`
	CreatedAt time.Time
	ExpireAt  time.Time `gorm:"DEFAULT: NULL"`
}

func (user *UserSession) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

type UserResult struct {
	ID        uint64 `gorm:"primary_key; AUTO_INCREMENT"`
	User      User   `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID    uint   `gorm:"not null"`
	Result    string `gorm:"not null"`
	CreatedAt time.Time
}

func (user *UserResult) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
