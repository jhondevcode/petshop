package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	FirstName string    `gorm:"not null" json:"firstName" validate:"required,min=5"`
	LastName  string    `gorm:"not null" json:"lastName" validate:"required,min=5"`
	Birthday  time.Time `gorm:"not null" json:"birthday"`
	Email     string    `gorm:"unique; not null" json:"email" validate:"required,email"`
	Gender    string    `gorm:"type:enum('V','M');size:1;not null" validate:"oneof=V M"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano" json:"updatedAt"`
	DeletedAt gorm.DeletedAt
	// DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *User) BeforeCreate(tx *gorm.DB) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

func (p *User) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

func (p *User) BeforeDelete(tx *gorm.DB) error {
	p.DeletedAt.Valid = true
	p.DeletedAt.Time = time.Now()
	return nil
}
