package models

type NextDataRes struct {
	Props struct {
		IsServer     bool `json:"isServer"`
		InitialState struct {
			Menu struct {
				MenuList []struct {
					ID         string `json:"id"`
					IsMainMenu bool   `json:"isMainMenu"`
					Name       string `json:"name"`
					IconText   string `json:"iconText"`
					Seo        struct {
						Slug        string `json:"slug"`
						URL         string `json:"url"`
						URLRedirect string `json:"url_redirect"`
						ShareURL    string `json:"share_url"`
					} `json:"seo"`
					Primary        bool   `json:"primary"`
					MenuURL        string `json:"menuUrl,omitempty"`
					Href           string `json:"href"`
					Slug           string `json:"slug"`
					DotCode        string `json:"dotCode"`
					Order          int    `json:"order"`
					DataSlug       string `json:"dataSlug,omitempty"`
					ParentID       string `json:"parentId"`
					HaveBanner     int    `json:"haveBanner"`
					TitleRibbon    string `json:"titleRibbon"`
					QuantityRibbon int    `json:"quantityRibbon"`
					SubMenu        []struct {
						ID         string `json:"id"`
						IsMainMenu bool   `json:"isMainMenu"`
						Name       string `json:"name"`
						IconText   string `json:"iconText"`
						Seo        struct {
							Slug        string `json:"slug"`
							URL         string `json:"url"`
							URLRedirect string `json:"url_redirect"`
							ShareURL    string `json:"share_url"`
						} `json:"seo"`
						Primary           bool          `json:"primary"`
						Href              string        `json:"href"`
						Slug              string        `json:"slug"`
						DotCode           string        `json:"dotCode"`
						Order             int           `json:"order"`
						DataSlug          string        `json:"dataSlug"`
						ParentID          string        `json:"parentId"`
						HaveBanner        int           `json:"haveBanner"`
						TitleRibbon       string        `json:"titleRibbon"`
						QuantityRibbon    int           `json:"quantityRibbon"`
						SubMenu           []interface{} `json:"subMenu"`
						Active            bool          `json:"active"`
						LayoutType        string        `json:"layoutType"`
						Highlights        []interface{} `json:"highlights"`
						URLRedirect       string        `json:"urlRedirect"`
						SubMenuRibbonData []interface{} `json:"subMenuRibbonData,omitempty"`
						SubMenuRibbon     []struct {
							ID                string `json:"id"`
							Name              string `json:"name"`
							NameFilterSmarttv string `json:"name_filter_smarttv"`
							Slug              string `json:"slug"`
							Seo               struct {
								Slug        string `json:"slug"`
								URL         string `json:"url"`
								URLRedirect string `json:"url_redirect"`
								ShareURL    string `json:"share_url"`
							} `json:"seo"`
							Images struct {
								Web struct {
									PosterV4          string `json:"poster_v4"`
									ThumbnailHotV4    string `json:"thumbnail_hot_v4"`
									ThumbnailBigV4    string `json:"thumbnail_big_v4"`
									ThumbnailV4       string `json:"thumbnail_v4"`
									CarouselTvV4      string `json:"carousel_tv_v4"`
									CarouselAppV4     string `json:"carousel_app_v4"`
									CarouselWebV4     string `json:"carousel_web_v4"`
									CarouselAppV4Ntc  string `json:"carousel_app_v4_ntc"`
									PosterV4Ntc       string `json:"poster_v4_ntc"`
									ThumbnailBigV4Ntc string `json:"thumbnail_big_v4_ntc"`
									ThumbnailHotV4Ntc string `json:"thumbnail_hot_v4_ntc"`
									ThumbnailV4Ntc    string `json:"thumbnail_v4_ntc"`
									CarouselWebV4Ntc  string `json:"carousel_web_v4_ntc"`
									CarouselTvV4Ntc   string `json:"carousel_tv_v4_ntc"`
									LivetvLogoLight   string `json:"livetv_logo_light"`
									LivetvLogoDark    string `json:"livetv_logo_dark"`
									LivetvLogoLight11 string `json:"livetv_logo_light_11"`
									LivetvLogoDark11  string `json:"livetv_logo_dark_11"`
									PosterV4Tablet    string `json:"poster_v4_tablet"`
									ThumbnailV4Tablet string `json:"thumbnail_v4_tablet"`
								} `json:"Web"`
								App struct {
									PosterV4          string `json:"poster_v4"`
									ThumbnailHotV4    string `json:"thumbnail_hot_v4"`
									ThumbnailBigV4    string `json:"thumbnail_big_v4"`
									ThumbnailV4       string `json:"thumbnail_v4"`
									CarouselTvV4      string `json:"carousel_tv_v4"`
									CarouselAppV4     string `json:"carousel_app_v4"`
									CarouselWebV4     string `json:"carousel_web_v4"`
									CarouselAppV4Ntc  string `json:"carousel_app_v4_ntc"`
									PosterV4Ntc       string `json:"poster_v4_ntc"`
									ThumbnailBigV4Ntc string `json:"thumbnail_big_v4_ntc"`
									ThumbnailHotV4Ntc string `json:"thumbnail_hot_v4_ntc"`
									ThumbnailV4Ntc    string `json:"thumbnail_v4_ntc"`
									CarouselWebV4Ntc  string `json:"carousel_web_v4_ntc"`
									CarouselTvV4Ntc   string `json:"carousel_tv_v4_ntc"`
									LivetvLogoLight   string `json:"livetv_logo_light"`
									LivetvLogoDark    string `json:"livetv_logo_dark"`
									LivetvLogoLight11 string `json:"livetv_logo_light_11"`
									LivetvLogoDark11  string `json:"livetv_logo_dark_11"`
									PosterV4Tablet    string `json:"poster_v4_tablet"`
									ThumbnailV4Tablet string `json:"thumbnail_v4_tablet"`
								} `json:"App"`
								Smarttv struct {
									PosterV4          string `json:"poster_v4"`
									ThumbnailHotV4    string `json:"thumbnail_hot_v4"`
									ThumbnailBigV4    string `json:"thumbnail_big_v4"`
									ThumbnailV4       string `json:"thumbnail_v4"`
									CarouselTvV4      string `json:"carousel_tv_v4"`
									CarouselAppV4     string `json:"carousel_app_v4"`
									CarouselWebV4     string `json:"carousel_web_v4"`
									CarouselAppV4Ntc  string `json:"carousel_app_v4_ntc"`
									PosterV4Ntc       string `json:"poster_v4_ntc"`
									ThumbnailBigV4Ntc string `json:"thumbnail_big_v4_ntc"`
									ThumbnailHotV4Ntc string `json:"thumbnail_hot_v4_ntc"`
									ThumbnailV4Ntc    string `json:"thumbnail_v4_ntc"`
									CarouselWebV4Ntc  string `json:"carousel_web_v4_ntc"`
									CarouselTvV4Ntc   string `json:"carousel_tv_v4_ntc"`
									LivetvLogoLight   string `json:"livetv_logo_light"`
									LivetvLogoDark    string `json:"livetv_logo_dark"`
									LivetvLogoLight11 string `json:"livetv_logo_light_11"`
									LivetvLogoDark11  string `json:"livetv_logo_dark_11"`
									PosterV4Tablet    string `json:"poster_v4_tablet"`
									ThumbnailV4Tablet string `json:"thumbnail_v4_tablet"`
								} `json:"Smarttv"`
								Tablet struct {
									PosterV4          string `json:"poster_v4"`
									ThumbnailHotV4    string `json:"thumbnail_hot_v4"`
									ThumbnailBigV4    string `json:"thumbnail_big_v4"`
									ThumbnailV4       string `json:"thumbnail_v4"`
									CarouselTvV4      string `json:"carousel_tv_v4"`
									CarouselAppV4     string `json:"carousel_app_v4"`
									CarouselWebV4     string `json:"carousel_web_v4"`
									CarouselAppV4Ntc  string `json:"carousel_app_v4_ntc"`
									PosterV4Ntc       string `json:"poster_v4_ntc"`
									ThumbnailBigV4Ntc string `json:"thumbnail_big_v4_ntc"`
									ThumbnailHotV4Ntc string `json:"thumbnail_hot_v4_ntc"`
									ThumbnailV4Ntc    string `json:"thumbnail_v4_ntc"`
									CarouselWebV4Ntc  string `json:"carousel_web_v4_ntc"`
									CarouselTvV4Ntc   string `json:"carousel_tv_v4_ntc"`
									LivetvLogoLight   string `json:"livetv_logo_light"`
									LivetvLogoDark    string `json:"livetv_logo_dark"`
									LivetvLogoLight11 string `json:"livetv_logo_light_11"`
									LivetvLogoDark11  string `json:"livetv_logo_dark_11"`
									PosterV4Tablet    string `json:"poster_v4_tablet"`
									ThumbnailV4Tablet string `json:"thumbnail_v4_tablet"`
								} `json:"Tablet"`
								Mobile struct {
									PosterV4          string `json:"poster_v4"`
									ThumbnailHotV4    string `json:"thumbnail_hot_v4"`
									ThumbnailBigV4    string `json:"thumbnail_big_v4"`
									ThumbnailV4       string `json:"thumbnail_v4"`
									CarouselTvV4      string `json:"carousel_tv_v4"`
									CarouselAppV4     string `json:"carousel_app_v4"`
									CarouselWebV4     string `json:"carousel_web_v4"`
									CarouselAppV4Ntc  string `json:"carousel_app_v4_ntc"`
									PosterV4Ntc       string `json:"poster_v4_ntc"`
									ThumbnailBigV4Ntc string `json:"thumbnail_big_v4_ntc"`
									ThumbnailHotV4Ntc string `json:"thumbnail_hot_v4_ntc"`
									ThumbnailV4Ntc    string `json:"thumbnail_v4_ntc"`
									CarouselWebV4Ntc  string `json:"carousel_web_v4_ntc"`
									CarouselTvV4Ntc   string `json:"carousel_tv_v4_ntc"`
									LivetvLogoLight   string `json:"livetv_logo_light"`
									LivetvLogoDark    string `json:"livetv_logo_dark"`
									LivetvLogoLight11 string `json:"livetv_logo_light_11"`
									LivetvLogoDark11  string `json:"livetv_logo_dark_11"`
									PosterV4Tablet    string `json:"poster_v4_tablet"`
									ThumbnailV4Tablet string `json:"thumbnail_v4_tablet"`
								} `json:"Mobile"`
							} `json:"images"`
							Thumbs interface{} `json:"thumbs"`
							Avatar string      `json:"avatar"`
						} `json:"subMenuRibbon,omitempty"`
					} `json:"subMenu"`
					Active      bool          `json:"active"`
					LayoutType  string        `json:"layoutType"`
					Highlights  []interface{} `json:"highlights"`
					URLRedirect string        `json:"urlRedirect"`
				} `json:"menuList"`
			} `json:"Menu"`
		} `json:"initialState"`
	} `json:"props"`
}
