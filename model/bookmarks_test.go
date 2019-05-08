package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sh-tatsuno/shabm/db"
	"github.com/sh-tatsuno/shabm/entity"
)

func TestGetAllBookmarks(t *testing.T) {

	t.Run("Normal", func(t *testing.T) {

		// ### Given ###
		expected := entity.Bookmark{
			URL:    "http://sample.com",
			Title:  "test",
			UserId: 1,
		}

		db.InitTestDB()
		defer db.CloseTestDB()

		dbc := db.Instance.DB.Create(&expected)
		if dbc.Error != nil {
			t.Fatalf("err should be nil. err: %v", dbc.Error)
		}

		b := BookmarkModel{}
		actual, err := b.GetAllBookmarks()
		if err != nil {
			t.Fatalf("err should be nil. err: %v", err)
		}

		if len(actual) > 1 {
			t.Fatalf("value lebgth should be 1. actual length: %v", len(actual))
		}

		if diff := cmp.Diff(actual[0], expected); diff != "" {
			t.Fatalf("Failed test in comparison. diff is below:\n %s\n", diff)
		}

	})

}
