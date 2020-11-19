package entity

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Model baseModel
type Model struct {
	ID        uint      `gorm:"column:id;primary_key;auto_increment;"`
	CreatedAt time.Time `gorm:"column:created_at;index;"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;"`
}

// TableName table name
func (Model) TableName(name string) string {
	return fmt.Sprintf("%s%s", viper.GetString("Gorm.TablePrefix"), name)
}
