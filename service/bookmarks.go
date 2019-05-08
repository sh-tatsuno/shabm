package service

import (
	"log"

	"github.com/sh-tatsuno/shabm/entity"
	"github.com/sh-tatsuno/shabm/model"
)

type BookmarkServiceInterface interface {
	GetAllBookmarks() ([]entity.Bookmark, error)
	CreateBookmark(url string, id uint) error
}

type BookmarkService struct {
	BookmarkModel    model.BookmarkModelInterface
	URLContentModel model.URLContentModelInterface
}

func NewBookmarkService() BookmarkService {
	return BookmarkService{
		BookmarkModel:    &model.BookmarkModel{},
		URLContentModel: &model.URLContentModel{},
	}
}

func (s *BookmarkService) GetAllBookmarks() ([]entity.Bookmark, error) {
	bookmarks, err := s.BookmarkModel.GetAllBookmarks()
	return bookmarks, err
}

func (s *BookmarkService) CreateBookmark(url string, id uint) error {

	title, err := s.URLContentModel.GetTitle(url)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	bookmark := entity.Bookmark{
		URL:    url,
		Title:  title,
		UserId: id,
	}

	err = s.BookmarkModel.CreateBookmark(bookmark)
	if err != nil {
		log.Printf("CreateBookmark error: %v", err)
		return err
	}
	return nil
}
