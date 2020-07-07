package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	DeletedAt time.Time `json:"deletedat" sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error during object creating %v", err)
	}

	err = scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("Error during object creating %v", err)
	}

	return nil
}
