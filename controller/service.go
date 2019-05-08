package controller

import (
	"github.com/sh-tatsuno/shabm/service"
)

type Services struct {
	PersonService  service.PersonServiceInterface
	BookmarkService service.BookmarkServiceInterface
}

func NewServices() Services {
	n1 := service.NewPersonService()
	n2 := service.NewBookmarkService()
	return Services{
		PersonService:  &n1,
		BookmarkService: &n2,
	}
}
