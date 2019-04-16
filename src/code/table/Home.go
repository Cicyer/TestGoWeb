package table

import "time"

type Home struct {
	ID         int       `gorm:"primary_key;"`
	Name       string    `gorm:"type:varchar(20);not null;"`
	CreateTime time.Time `gorm:"type:datetime;DEFAULT NULL;"`
}

func (Home) TableName() string {
	return "home"
}
