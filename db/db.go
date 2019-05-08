package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // use sqlite
	"github.com/sh-tatsuno/shabm/entity"
)

type ORM struct {
	DB *gorm.DB
	Tx *gorm.DB
}

// Instance : singleton
var Instance ORM

func InitDB() {
	if (Instance == ORM{}) {
		var err error
		Instance.DB, err = gorm.Open("sqlite3", "test.sqlite3")
		if err != nil {
			s := fmt.Sprintf("failed to connect database. err: %v\n", err)
			panic(s)
		}
	}

	Instance.DB.LogMode(true)

	Instance.DB.AutoMigrate(
		&entity.Person{},
		&entity.Bookmark{},
	)
}

func InitTestDB() {
	if (Instance == ORM{}) {
		var err error
		Instance.DB, err = gorm.Open("sqlite3", "testy.sqlite3")
		if err != nil {
			s := fmt.Sprintf("failed to connect database. err: %v\n", err)
			panic(s)
		}
	}

	Instance.DB.LogMode(true)

	Instance.DB.AutoMigrate(
		&entity.Person{},
		&entity.Bookmark{},
	)

	Instance.Tx = Instance.DB.Begin()
}

func CloseTestDB() {
	Instance.DB.Close()
	path := "testy.sqlite3"
	if f, err := os.Stat(path); os.IsNotExist(err) || f.IsDir() {
		fmt.Println("no exist file")
	} else {
		if err := os.Remove(path); err != nil {
			fmt.Println(err)
		}
	}
}

func TransactionHandlerFunc() gin.HandlerFunc {

	return func(c *gin.Context) {
		Instance.Tx = Instance.DB.Begin()
		if Instance.Tx.Error != nil {
			log.Printf("error: %v", Instance.Tx.Error)
			return
		}
		defer func() {
			if len(c.Errors) > 0 {
				Instance.Tx.Rollback()
				log.Println("END TRANSACTION ROLLBACK")
				return
			}
			Instance.Tx.Commit()
			log.Println("END TRANSACTION COMMIT")
		}()

		log.Println("START TRANSACTION")
		c.Next()
	}
}
