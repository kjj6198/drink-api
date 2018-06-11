package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DataBaseOptions struct {
	Addr     string
	Port     string
	User     string
	Database string
	Password string
	SSLMode  bool
}

func (option *DataBaseOptions) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=\"%s\" sslmode=disable",
		option.Addr,
		option.Port,
		option.User,
		option.Database,
		option.Password,
	)
}

// Connect connect db
func Connect(option *DataBaseOptions) (db *gorm.DB, err error) {
	return gorm.Open("postgres", option.String())
}

func Close(db *gorm.DB) {
	db.Close()
}
