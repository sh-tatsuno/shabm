package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sh-tatsuno/shabm/service"
)

func HomeIndex(
	p service.PersonServiceInterface,
	b service.BookmarkServiceInterface,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		people, err := p.GetAllPeople()
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
			return
		}
		bookmarks, err := b.GetAllBookmarks()
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
			return
		}

		c.HTML(200, "index.tmpl", gin.H{
			"people": people,
			"bkm":    bookmarks,
		})
	}
}
