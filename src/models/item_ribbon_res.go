package models

type ItemRibbonRes struct {
	ID               string `json:"id" bson:"id"`
	GroupID          string `json:"group_id" bson:"group_id"`
	GeoCheck         int    `json:"geo_check" bson:"geo_check"`
	Type             int    `json:"type" bson:"type"`
	Category         int    `json:"category" bson:"category"`
	Title            string `json:"title" bson:"title"`
	ShortDescription string `json:"short_description" bson:"short_description"`
	Genre            string `json:"genre" bson:"genre"`
	Resolution       int    `json:"resolution" bson:"resolution"`
	Images           struct {
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
	AvgRate        float64 `json:"avg_rate" bson:"avg_rate"`
	TotalRate      int     `json:"total_rate" bson:"total_rate"`
	IsWatchlater   bool    `json:"is_watchlater" bson:"is_watchlater"`
	IsPremium      int     `json:"is_premium" bson:"is_premium"`
	IsNew          int     `json:"is_new" bson:"is_new"`
	Episode        int     `json:"episode" bson:"episode"`
	CurrentEpisode string  `json:"current_episode" bson:"current_episode"`
	Seo            struct {
		Slug        string `json:"slug" bson:"slug"`
		URL         string `json:"url" bson:"url"`
		URLRedirect string `json:"url_redirect" bson:"url_redirect"`
		ShareURL    string `json:"share_url" bson:"share_url"`
	} `json:"seo" bson:"seo"`
	Slug               string `json:"slug" bson:"slug"`
	DisplayImageType   int    `json:"display_image_type" bson:"display_image_type"`
	LabelSubtitleAudio string `json:"label_subtitle_audio" bson:"label_subtitle_audio"`
	LabelPublicDay     string `json:"label_public_day" bson:"label_public_day"`
	LinkPlay           struct {
		HlsLinkPlay    string `json:"hls_link_play" bson:"hls_link_play"`
		PosterLinkPlay string `json:"poster_link_play" bson:"poster_link_play"`
	} `json:"link_play" bson:"link_play"`
	TagsDisplay   []interface{} `json:"tags_display" bson:"tags_display"`
	Ranking       int           `json:"ranking" bson:"ranking"`
	Runtime       int           `json:"runtime" bson:"runtime"`
	ExternalURL   string        `json:"external_url" bson:"external_url"`
	AllowClick    int           `json:"allow_click" bson:"allow_click"`
	SeasonName    string        `json:"season_name" bson:"season_name"`
	RelatedSeason []interface{} `json:"related_season" bson:"related_season"`
	AgeRange      string        `json:"age_range" bson:"age_range"`
	WarningTag    string        `json:"warning_tag" bson:"warning_tag"`
	VodSchedule   struct {
		Day     string `json:"day" bson:"day"`
		Channel string `json:"channel" bson:"channel"`
		Hour    string `json:"hour" bson:"hour"`
	} `json:"vod_schedule" bson:"vod_schedule"`
	IsFavorite            bool    `json:"is_favorite" bson:"is_favorite"`
	IsComingSoon          int     `json:"is_coming_soon" bson:"is_coming_soon"`
	IsLive                int     `json:"is_live" bson:"is_live"`
	IsFinish              int     `json:"is_finish" bson:"is_finish"`
	StartTime             int     `json:"start_time" bson:"start_time"`
	EndTime               int     `json:"end_time" bson:"end_time"`
	DeepLink              string  `json:"deep_link" bson:"deep_link"`
	OptionDirect          string  `json:"option_direct" bson:"option_direct"`
	AwardDescription      string  `json:"award_description" bson:"award_description"`
	IsPremiumDisplay      string  `json:"is_premium_display" bson:"is_premium_display"`
	PackageDenyContent    []int   `json:"package_deny_content" bson:"package_deny_content"`
	OptionDirectPackageID int     `json:"option_direct_package_id" bson:"option_direct_package_id"`
	Odr                   float64 `json:"odr" bson:"odr"`
	PeopleStr             string  `json:"people_str" bson:"people_str"`
	TagsGenreTxt          string  `json:"tags_genre_txt" bson:"tags_genre_txt"`
	IsPremiere            int     `json:"is_premiere" bson:"is_premiere"`
	Tvod                  struct {
	} `json:"tvod" bson:"tvod"`
	IsSimulcast        int    `json:"is_simulcast" bson:"is_simulcast"`
	HasObjectDetection bool   `json:"has_object_detection" bson:"has_object_detection"`
	CategoryID         string `json:"category_id" bson:"category_id"`
}
