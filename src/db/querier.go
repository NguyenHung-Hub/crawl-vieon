package db

import (
	"context"
	"crawl/src/models"
)

type Querier interface {
	CreateUrlCrawl(ctx context.Context, data []*models.UrlCrawl) error
	CreateRibbonIDCrawl(ctx context.Context, data []*models.RibbonIDCrawl) error
	CreateContentIDCrawl(ctx context.Context, data []*models.ContentIDCrawl) error

	GetAllUrl(ctx context.Context) (*[]models.UrlCrawl, error)
	GetAllRibbonID(ctx context.Context) (*[]models.OnlyRibbonID, error)
	GetAllContentID(ctx context.Context) (*[]models.OnlyContentID, error)

	IsExistsRibbon(ctx context.Context, id string) bool
	CreateRibbon(ctx context.Context, ribbon *models.Ribbon) error
	CreateContent(ctx context.Context, content *models.Content) error
}

var _ Querier = (*Queries)(nil)
