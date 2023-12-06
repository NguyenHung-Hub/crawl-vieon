package db

import (
	"context"
	"crawl/src/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

type OnlyContentIDVieon struct {
	ID string `json:"id"`
}

func (q *Queries) CreateUrlCrawl(ctx context.Context, data []*models.UrlCrawl) error {
	err := q.db.Create(data).Error
	return err
}
func (q *Queries) CreateRibbonIDCrawl(ctx context.Context, data []*models.RibbonIDCrawl) error {
	err := q.db.Create(data).Error
	return err
}
func (q *Queries) CreateContentIDCrawl(ctx context.Context, data []*models.ContentIDCrawl) error {
	err := q.db.Create(data).Error
	return err
}

func (q *Queries) GetAllUrl(ctx context.Context) (*[]models.UrlCrawl, error) {
	var data []models.UrlCrawl
	err := q.db.Find(&data).Error
	return &data, err
}
func (q *Queries) GetAllRibbonID(ctx context.Context) (*[]models.OnlyRibbonID, error) {
	var data []models.OnlyRibbonID
	err := q.db.Find(&data).Error
	return &data, err
}
func (q *Queries) GetAllContentID(ctx context.Context) (*[]models.OnlyContentID, error) {
	var data []models.OnlyContentID
	err := q.db.Find(&data).Error
	return &data, err
}

func (q *Queries) IsExistsRibbon(ctx context.Context, id string) bool {

	var ribbon models.Ribbon
	err := q.db.Where("ribbon_id = ?", id).First(&ribbon).Error
	if err != nil {
		log.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		} else {
			return true
		}

	}

	if ribbon.RibbonID != "" && ribbon.Name != "" {
		return true
	}

	return false
}

func (q *Queries) CreateRibbon(ctx context.Context, ribbon *models.Ribbon) error {

	err := q.db.Create(ribbon).Error
	return err
}

func (q *Queries) CreateContent(ctx context.Context, content *models.Content) error {
	err := q.db.Create(content).Error
	return err
}
