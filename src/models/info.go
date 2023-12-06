package models

import (
	"gorm.io/gorm"
)

type UrlCrawl struct {
	gorm.Model
	Url string `json:"url" gorm:"url;"`
}
type RibbonIDCrawl struct {
	gorm.Model
	RibbonID string `json:"ribbon_id" gorm:"ribbon_id;"`
}
type ContentIDCrawl struct {
	gorm.Model
	ContentID string `json:"content_id" gorm:"content_id;"`
}

func (UrlCrawl) TableName() string       { return "url_crawl" }
func (RibbonIDCrawl) TableName() string  { return "ribbon_id_crawl" }
func (ContentIDCrawl) TableName() string { return "content_id_crawl" }
