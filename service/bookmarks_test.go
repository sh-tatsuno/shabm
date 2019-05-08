package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/sh-tatsuno/shabm/entity"
	"github.com/sh-tatsuno/shabm/model/mock_model"
)

func TestGetAllBookmarks(t *testing.T) {

	t.Run("Normal", func(t *testing.T) {

		// ### Given ###
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockBookmarkModel := mock_model.NewMockBookmarkModelInterface(mockCtrl)
		expected := []entity.Bookmark{
			entity.Bookmark{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: nil,
				},
				URL:    "http://sample.com",
				Title:  "test",
				UserId: 1,
			},
		}
		mockBookmarkModel.EXPECT().GetAllBookmarks().Return(expected, nil)

		mockURLContentModel := mock_model.NewMockURLContentModelInterface(mockCtrl)

		s := BookmarkService{
			BookmarkModel:   mockBookmarkModel,
			URLContentModel: mockURLContentModel,
		}

		// ### When ###
		actual, err := s.GetAllBookmarks()

		// ### Then ###
		if err != nil {
			t.Fatalf("Failed test.err should be nil. err: %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Failed test. expected: %v, but actual: %v", expected, actual)
		}

	})

}
