package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint64       `gorm:"id;primaryKey"`
	Name      string       `gorm:"name"`
	CreatedAt time.Time    `gorm:"created_at"`
	UpdatedAt time.Time    `gorm:"updated_at"`
	DeletedAt sql.NullTime `gorm:"deleted_at"`
}

func (*User) TableName() string {
	return "users"
}
