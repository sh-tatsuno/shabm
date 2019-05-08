package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/xerrors"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/sh-tatsuno/shabm/entity"
	"github.com/sh-tatsuno/shabm/service/mock_service"
)

func TestHome(t *testing.T) {

	// ### Setting ###
	gin.SetMode(gin.TestMode)

	t.Run("Normal: 200", func(t *testing.T) {

		// ### Given ###
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockPersonService := mock_service.NewMockPersonServiceInterface(mockCtrl)
		p := []entity.Person{
			entity.Person{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: nil,
				},
				Name: "Taro",
				Age:  10,
			},
		}
		mockPersonService.EXPECT().GetAllPeople().Return(p, nil)

		mockCtrl = gomock.NewController(t)
		mockBookmarkService := mock_service.NewMockBookmarkServiceInterface(mockCtrl)
		b := []entity.Bookmark{
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
		mockBookmarkService.EXPECT().GetAllBookmarks().Return(b, nil)

		r := gin.New()
		r.LoadHTMLGlob("../../view/*")
		r.GET("/", HomeIndex(mockPersonService, mockBookmarkService))

		// ### When ###
		req, _ := http.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// ### Then ###
		expected := http.StatusOK
		actual := w.Code
		if actual != expected {
			t.Fatalf("Failed test. expected: %v, but actual: %v", expected, actual)
		}
	})

	t.Run("Normal (empty data): 200", func(t *testing.T) {

		// ### Given ###
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockPersonService := mock_service.NewMockPersonServiceInterface(mockCtrl)
		p := []entity.Person{}
		mockPersonService.EXPECT().GetAllPeople().Return(p, nil)

		mockCtrl = gomock.NewController(t)
		mockBookmarkService := mock_service.NewMockBookmarkServiceInterface(mockCtrl)
		b := []entity.Bookmark{}
		mockBookmarkService.EXPECT().GetAllBookmarks().Return(b, nil)

		r := gin.New()
		r.LoadHTMLGlob("../../view/*")
		r.GET("/", HomeIndex(mockPersonService, mockBookmarkService))

		// ### When ###
		req, _ := http.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// ### Then ###
		expected := http.StatusOK
		actual := w.Code
		if actual != expected {
			t.Fatalf("Failed test. expected: %v, but actual: %v", expected, actual)
		}
	})

	t.Run("Abnormal (error in PersonService): 500", func(t *testing.T) {

		// ### Given ###
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockPersonService := mock_service.NewMockPersonServiceInterface(mockCtrl)
		p := []entity.Person{}
		err := xerrors.New("Error in Person")
		mockPersonService.EXPECT().GetAllPeople().Return(p, err)

		mockCtrl = gomock.NewController(t)
		mockBookmarkService := mock_service.NewMockBookmarkServiceInterface(mockCtrl)
		b := []entity.Bookmark{}
		mockBookmarkService.EXPECT().GetAllBookmarks().Return(b, nil)

		r := gin.New()
		r.LoadHTMLGlob("../../view/*")
		r.GET("/", HomeIndex(mockPersonService, mockBookmarkService))

		// ### When ###
		req, _ := http.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// ### Then ###
		expected := http.StatusInternalServerError
		actual := w.Code
		if actual != expected {
			t.Fatalf("Failed test. expected: %v, but actual: %v", expected, actual)
		}
	})

	t.Run("Abnormal (error in BookmarkService): 500", func(t *testing.T) {

		// ### Given ###
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockPersonService := mock_service.NewMockPersonServiceInterface(mockCtrl)
		p := []entity.Person{}
		mockPersonService.EXPECT().GetAllPeople().Return(p, nil)

		mockCtrl = gomock.NewController(t)
		mockBookmarkService := mock_service.NewMockBookmarkServiceInterface(mockCtrl)
		b := []entity.Bookmark{}
		err := xerrors.New("Error in Bookmark")
		mockBookmarkService.EXPECT().GetAllBookmarks().Return(b, err)

		r := gin.New()
		r.LoadHTMLGlob("../../view/*")
		r.GET("/", HomeIndex(mockPersonService, mockBookmarkService))

		// ### When ###
		req, _ := http.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// ### Then ###
		expected := http.StatusInternalServerError
		actual := w.Code
		if actual != expected {
			t.Fatalf("Failed test. expected: %v, but actual: %v", expected, actual)
		}
	})

}
