package controller

import (
	_ "net/http" // gin

	// _ "github.com/mattn/go-sqlite3" // db init

	"github.com/gin-gonic/gin"
	"github.com/sh-tatsuno/shabm/controller/handler"
	"github.com/sh-tatsuno/shabm/db"
)

func WebServer() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	db.InitDB()
	defer db.Instance.DB.Close()
	r.Use(db.TransactionHandlerFunc())

	s := NewServices()

	r.GET("/", handler.HomeIndex(s.PersonService, s.BookmarkService))
	r.POST("/form", handler.FormNew(s.PersonService, s.BookmarkService))

	r.Run()
}
