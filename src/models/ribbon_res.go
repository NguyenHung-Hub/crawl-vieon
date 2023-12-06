package models

type RibbonRes struct {
	ID               string `json:"id" bson:"category_id"`
	Name             string `json:"name" bson:"name"`
	Type             int64  `json:"type" bson:"type"`
	DisplayImageType int64  `json:"display_image_type" bson:"display_image_type"`
	IsPremium        int64  `json:"is_premium" bson:"is_premium"`
	Menus            []struct {
		ID string `json:"id" bson:"id"`
	} `json:"menus" bson:"menus"`
	Tags         []interface{}   `json:"tags" bson:"tags"`
	Items        []ItemRibbonRes `json:"items" bson:"-"`
	TrackingData struct {
		RecommendationID string `json:"recommendation_id" bson:"recommendation_id"`
		Type             string `json:"type" bson:"type"`
	} `json:"tracking_data" bson:"tracking_data"`
	Seo struct {
		Slug        string `json:"slug" bson:"slug"`
		URL         string `json:"url" bson:"url"`
		URLRedirect string `json:"url_redirect" bson:"url_redirect"`
		ShareURL    string `json:"share_url" bson:"share_url"`
	} `json:"seo" bson:"seo"`
	GeoCheck   int64 `json:"geo_check" bson:"geo_check"`
	Properties struct {
		Line      int64 `json:"line" bson:"line"`
		IsTitle   int64 `json:"is_title" bson:"is_title"`
		IsSlide   int64 `json:"is_slide" bson:"is_slide"`
		IsRefresh int64 `json:"is_refresh" bson:"is_refresh"`
		IsViewAll int64 `json:"is_view_all" bson:"is_view_all"`
		Thumb     int64 `json:"thumb" bson:"thumb"`
	} `json:"properties" bson:"properties"`
	Images struct {
		ThumbnailHotV4    string `json:"thumbnail_hot_v4" bson:"thumbnail_hot_v4"`
		ThumbnailBigV4    string `json:"thumbnail_big_v4" bson:"thumbnail_big_v4"`
		CarouselTvV4      string `json:"carousel_tv_v4" bson:"carousel_tv_v4"`
		CarouselAppV4     string `json:"carousel_app_v4" bson:"carousel_app_v4"`
		CarouselWebV4     string `json:"carousel_web_v4" bson:"carousel_web_v4"`
		ThumbnailV4       string `json:"thumbnail_v4" bson:"thumbnail_v4"`
		PosterV4          string `json:"poster_v4" bson:"poster_v4"`
		PromotionBanner   string `json:"promotion_banner" bson:"promotion_banner"`
		TitleCardLight    string `json:"title_card_light" bson:"title_card_light"`
		TitleCardDark     string `json:"title_card_dark" bson:"title_card_dark"`
		PosterOriginal    string `json:"poster_original" bson:"poster_original"`
		ThumbOriginal     string `json:"thumb_original" bson:"thumb_original"`
		ThumbnailHotV4Ntc string `json:"thumbnail_hot_v4_ntc" bson:"thumbnail_hot_v4_ntc"`
		ThumbnailBigV4Ntc string `json:"thumbnail_big_v4_ntc" bson:"thumbnail_big_v4_ntc"`
		CarouselTvV4Ntc   string `json:"carousel_tv_v4_ntc" bson:"carousel_tv_v4_ntc"`
		CarouselAppV4Ntc  string `json:"carousel_app_v4_ntc" bson:"carousel_app_v4_ntc"`
		CarouselWebV4Ntc  string `json:"carousel_web_v4_ntc" bson:"carousel_web_v4_ntc"`
		ThumbnailV4Ntc    string `json:"thumbnail_v4_ntc" bson:"thumbnail_v4_ntc"`
		PosterV4Ntc       string `json:"poster_v4_ntc" bson:"poster_v4_ntc"`
		LivetvLogoLight   string `json:"livetv_logo_light" bson:"livetv_logo_light"`
		LivetvLogoDark    string `json:"livetv_logo_dark" bson:"livetv_logo_dark"`
		LivetvLogoLight11 string `json:"livetv_logo_light_11" bson:"livetv_logo_light_11"`
		LivetvLogoDark11  string `json:"livetv_logo_dark_11" bson:"livetv_logo_dark_11"`
		RibbonDetailApp   string `json:"ribbon_detail_app" bson:"ribbon_detail_app"`
		RibbonDetailWeb   string `json:"ribbon_detail_web" bson:"ribbon_detail_web"`
		PromotionBannerSm string `json:"promotion_banner_sm" bson:"promotion_banner_sm"`
		PosterV4Original  string `json:"poster_v4_original" bson:"poster_v4_original"`
		PosterV4Tablet    string `json:"poster_v4_tablet" bson:"poster_v4_tablet"`
		ThumbnailV4Tablet string `json:"thumbnail_v4_tablet" bson:"thumbnail_v4_tablet"`
	} `json:"images" bson:"images"`
	Description   string `json:"description" bson:"description"`
	RecommendType string `json:"recommend_type" bson:"recommend_type"`
	DisplayFor    int64  `json:"display_for" bson:"display_for"`
	Odr           int64  `json:"odr" bson:"odr"`
	DurationTrans int64  `json:"duration_trans" bson:"duration_trans"`
	IsItems       int64  `json:"is_items" bson:"is_items"`
	AgeRange      string `json:"age_range" bson:"age_range"`
	Metadata      struct {
		Total int64 `json:"total" bson:"total"`
		Limit int64 `json:"limit" bson:"limit"`
		Page  int64 `json:"page" bson:"page"`
	} `json:"metadata" bson:"metadata"`
}
