package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sh-tatsuno/shabm/entity"
	"github.com/sh-tatsuno/shabm/service"
)

func FormNew(
	p service.PersonServiceInterface,
	b service.BookmarkServiceInterface,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))
		person := entity.Person{
			Name: name,
			Age:  age,
		}
		id, err := p.CreatePerson(person)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
			return
		}

		url := c.PostForm("url")
		err = b.CreateBookmark(url, id)
		if err != nil {
			log.Printf("error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
			return
		}

		c.Redirect(302, "/")
	}
}
