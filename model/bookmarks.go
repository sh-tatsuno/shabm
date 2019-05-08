package model

import (
	"github.com/sh-tatsuno/shabm/db"
	"github.com/sh-tatsuno/shabm/entity"
)

type BookmarkModelInterface interface {
	GetAllBookmarks() ([]entity.Bookmark, error)
	CreateBookmark(bookmark entity.Bookmark) error
}

type BookmarkModel struct {
}

func (m *BookmarkModel) GetAllBookmarks() ([]entity.Bookmark, error) {
	var bookmarks []entity.Bookmark
	dbc := db.Instance.Tx.Find(&bookmarks)
	return bookmarks, dbc.Error
}

func (m *BookmarkModel) CreateBookmark(bookmark entity.Bookmark) error {
	dbc := db.Instance.Tx.Create(&bookmark)
	return dbc.Error
}
