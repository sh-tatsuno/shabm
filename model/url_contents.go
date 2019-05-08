package model

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

type URLContentModelInterface interface {
	GetTitle(url string) (string, error)
}

type URLContentModel struct {
}

func (u *URLContentModel) GetTitle(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Printf("func=NewURLBody err=%s\n", err.Error())
		return "", err
	}
	return doc.Find("title").Text(), err
}
