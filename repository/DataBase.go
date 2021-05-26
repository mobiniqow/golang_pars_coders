package repository

import (
	"log"
	"mobiniqow/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// singlethon
type Server struct {
	DB *gorm.DB
}

var instantiated *Server = nil

func (c *Server) Connect() *Server {
	if instantiated == nil {
		var err error
		instantiated = &Server{}
		DBURL := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=postgres"
		instantiated.DB, err = gorm.Open("postgres", DBURL)
		if err != nil {
			log.Fatal("This is the error:", err)
		} else {
		}
		instantiated.DB.Debug().AutoMigrate(&model.Coin{}, &model.Wallet{})
	}
	return instantiated
}
