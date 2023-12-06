package models

import "gorm.io/gorm"

type Ribbon struct {
	gorm.Model
	RibbonID          string `json:"id" gorm:"column:ribbon_id;size:43;index"`
	Name              string `json:"name" gorm:"column:name"`
	Type              int64  `json:"type" gorm:"column:type"`
	DisplayImageType  int64  `json:"display_image_type" gorm:"column:display_image_type"`
	IsPremium         int64  `json:"is_premium" gorm:"column:is_premium"`
	ThumbnailHotV4    string `json:"thumbnail_hot_v4" gorm:"column:thumbnail_hot_v4"`
	ThumbnailBigV4    string `json:"thumbnail_big_v4" gorm:"column:thumbnail_big_v4"`
	ThumbnailV4       string `json:"thumbnail_v4" gorm:"column:thumbnail_v4"`
	ThumbnailHotV4Ntc string `json:"thumbnail_hot_v4_ntc" gorm:"column:thumbnail_hot_v4_ntc"`
	ThumbnailBigV4Ntc string `json:"thumbnail_big_v4_ntc" gorm:"column:thumbnail_big_v4_ntc"`
	ThumbnailV4Ntc    string `json:"thumbnail_v4_ntc" gorm:"column:thumbnail_v4_ntc"`
}

type OnlyRibbonID struct {
	RibbonID string `json:"id" gorm:"column:ribbon_id;"`
}

func (Ribbon) TableName() string       { return "ribbons" }
func (OnlyRibbonID) TableName() string { return Ribbon{}.TableName() }
